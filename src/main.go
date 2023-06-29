package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func getUuid() string {
	b := make([]byte, 16)
	_, _ = io.ReadFull(rand.Reader, b)

	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func main() {
	var elfPath string
	var libcPath string

	flag.StringVar(&elfPath, "elf", "/elf", "The path of ELF you want to run.")
	flag.StringVar(&libcPath, "libc", "/libc.so.6", "The path of Libc you want to use.")
	flag.Parse()
	fmt.Printf("Vessel: elfPath \t%s\n", elfPath)
	fmt.Printf("Vessel: libcPath \t%s\n", libcPath)
	if !isExist(elfPath) {
		fmt.Println("Vessel: elfPath not exist!")
		os.Exit(1)
	}
	if !isExist(libcPath) {
		fmt.Println("Vessel: libcPath not exist!")
		os.Exit(1)
	}

	var nameSpaceFlag int
	nameSpaceFlag |= syscall.CLONE_NEWNS

	nameSpaceFlag |= syscall.CLONE_NEWUTS

	nameSpaceFlag |= syscall.CLONE_NEWIPC

	nameSpaceFlag |= syscall.CLONE_NEWNET

	nameSpaceFlag |= syscall.CLONE_NEWPID

	//nameSpaceFlag |= syscall.CLONE_NEWUSER

	// 创建新的 namespace
	if err := syscall.Unshare(nameSpaceFlag); err != nil {
		fmt.Printf("Vessel: failed to create new mount namespace: %v\n", err)
		os.Exit(1)
	}

	if err := syscall.Sethostname([]byte("Vessel")); err != nil {
		fmt.Printf("Vessel: failed to set hostname: %v\n", err)
		os.Exit(1)
	}

	// 在新的 mount namespace 中挂载/proc目录
	if err := syscall.Mount("proc", "/proc", "proc", uintptr(syscall.MS_NOEXEC|syscall.MS_NOSUID|syscall.MS_NODEV), ""); err != nil {
		fmt.Printf("Vessel: failed to mount procfs: %v\n", err)
		os.Exit(1)
	}

	// 设置 LD_LIBRARY_PATH 环境变量，指定新的 glibc 库路径
	if err := os.Setenv("LD_PRELOAD", libcPath); err != nil {
		fmt.Printf("Vessel: failed to set LD_LIBRARY_PATH environment variable: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(elfPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// 隔离 uts,ipc,pid,mount,user,network
		Cloneflags: syscall.CLONE_NEWUSER,
		// 设置容器的UID和GID
		UidMappings: []syscall.SysProcIDMap{
			{
				// 容器的UID
				ContainerID: 1,
				// 宿主机的UID
				HostID: 0,
				Size:   1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				// 容器的GID
				ContainerID: 1,
				// 宿主机的GID
				HostID: 0,
				Size:   1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println()
		log.Fatalf("Vessel: %s\n", err)
	}
	// 卸载/proc目录
	if err := syscall.Unmount("/proc", syscall.MNT_DETACH); err != nil {
		fmt.Printf("Vessel: failed to unmount procfs: %v\n", err)
		os.Exit(1)
	}
}
