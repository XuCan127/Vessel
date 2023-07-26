package subSystems


import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"syscall"
)

func SetCpuCGroup(cpuPath string, pid int) error {

	//
	if err := os.MkdirAll(cpuPath, 0700); err != nil {
		return fmt.Errorf("create cgroup path fail err=%s", err)
	}
	if err := ioutil.WriteFile(path.Join(cpuPath, "cpu.cfs_quota_us"), []byte("10000"), 0700); err != nil {
		return fmt.Errorf("write cpu quota us fail err=%s", err)
	}
	if err := ioutil.WriteFile(path.Join(cpuPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
		return fmt.Errorf("write cpu tasks  fail err=%s", err)
	}
	return nil
}
func CleanCpuCGroup(cpuPath string) error {

	if err := syscall.Rmdir(cpuPath); err != nil {
		panic(err)
	}
	return nil

}
