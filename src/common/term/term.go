package term

import (
	"Vessel/src/server/constant"
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

/*
	标准化工具类，用于屏蔽系统差异
*/

const (
	Perm0777 = 0777 // 用户、组用户和其它用户都有读/写/执行权限
	Perm0755 = 0755 // 用户具有读/写/执行权限，组用户和其它用户具有读/写权限；
	Perm0644 = 0644 // 用户具有读/写权限，组用户和其它用户具只读权限；
	Perm0622 = 0622 // 用户具有读/写权限，组用户和其它用户具只写权限；
)

const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z07:00"
const CreateTimeFmt = "01月02日 15时04分"

const ErrorFmt = "failed to %s: %v.\n"

type TcpConfig struct {
	IP   []byte
	Port int
}

func StdStreams() (stdIn io.ReadCloser, stdOut, stdErr io.Writer) {
	return os.Stdin, os.Stdout, os.Stderr
}

func CheckFile(path string, defaultContent []byte) {
	// 检查文件是否存在
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 文件不存在，创建文件并写入内容
		err := os.WriteFile(path, defaultContent, 0644)
		if err != nil {
			fmt.Println("无法创建文件:", err)
			return
		}

		fmt.Println("文件已创建并写入内容")
	} else {
		// 其他错误情况
		fmt.Println("发生错误:", err)
	}
}

func CheckFS(filesystemName string) bool {
	data, err := os.ReadFile("/proc/filesystems")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Contains(string(data), filesystemName)
}

// LoopBack 内网回路
func LoopBack(port int) TcpConfig {
	return TcpConfig{IP: net.IPv4(127, 0, 0, 1), Port: port}
}

// Wildcard 外网回路
func Wildcard(port int) TcpConfig {
	return TcpConfig{IP: net.IPv4(0, 0, 0, 0), Port: port}
}

func NewPipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return read, write, nil
}

func PivotRoot(root string) error {
	if err := Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf(ErrorFmt, "mount", err)
	}

	pivotDir := filepath.Join(root, ".pivot_root")
	if err := os.Mkdir(pivotDir, 0777); err != nil {
		return err
	}

	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		return fmt.Errorf(ErrorFmt, "pivot_root", err)
	}

	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf(ErrorFmt, "chdir", err)
	}

	pivotDir = filepath.Join("/", ".pivot_root")

	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf(ErrorFmt, "unmount pivot_root", err)
	}

	return os.Remove(pivotDir)
}

func Unzip(zipPath, destPath string) error {
	// 打开ZIP文件
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	var errorCount int // 错误计数器

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			errorCount++
			continue
		}

		targetPath := filepath.Join(destPath, f.Name)

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(targetPath, f.Mode())
			if err != nil {
				errorCount++
				rc.Close()
				continue
			}
		} else {
			err = os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
			if err != nil {
				errorCount++
				rc.Close()
				continue
			}

			targetFile, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				errorCount++
				rc.Close()
				continue
			}

			_, err = io.Copy(targetFile, rc)
			targetFile.Close()
			if err != nil {
				errorCount++
				continue
			}
		}

		rc.Close()
	}

	if errorCount > constant.Retry {
		// 出现错误，删除已解压的部分
		os.RemoveAll(destPath)
		return fmt.Errorf("unzip failed with %d errors", errorCount)
	}

	return nil
}
