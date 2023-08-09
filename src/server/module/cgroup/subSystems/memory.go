package subSystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"syscall"
)

const (
	memoryLimit = "memory.limit_in_bytes"
)

func SetMemoryCGroup(memoryPath string, pid int, res *ResourceConfig) error {
	if err := os.MkdirAll(memoryPath, 0700); err != nil {
		return fmt.Errorf("create cgroup path fail err=%s", err)
	}
	if err := ioutil.WriteFile(path.Join(memoryPath, memoryLimit), []byte(res.Memory), 0700); err != nil {
		return fmt.Errorf("write memory limit bytes fail err=%s", err)
	}
	if err := ioutil.WriteFile(path.Join(memoryPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("write memory tasks  fail err=%s", err)
	}

	return nil
}

func CleanMemoryCGroup(memoryPath string) error {
	if err := syscall.Rmdir(memoryPath); err != nil {
		panic(err)
	}
	return nil
}
