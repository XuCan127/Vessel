package term

import "syscall"

func mount(source string, target string, fstype string, flags uintptr, data string) (err error) {
	return syscall.Mount(source, target, fstype, flags, data)
}

/*
	对进程操作的工具类
*/

func IsProcessAlive(pid int) bool {
	err := syscall.Kill(pid, syscall.Signal(0))
	if err == nil || err == syscall.EPERM {
		return true
	}

	return false
}

func KillProcess(pid int) {
	syscall.Kill(pid, syscall.SIGKILL)
}
