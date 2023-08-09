package term

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
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

const ErrorFmt = "failed to %s: %v.\n"

type TcpConfig struct {
	IP   []byte
	Port int
}

func StdStreams() (stdIn io.ReadCloser, stdOut, stdErr io.Writer) {
	return os.Stdin, os.Stdout, os.Stderr
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
	if err := mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
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
