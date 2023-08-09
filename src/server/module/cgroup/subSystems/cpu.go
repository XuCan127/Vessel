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
	cpu = "cpu,cpuacct"
	// 表示一个cpu带宽，单位为微秒。系统总CPU带宽： cpu核心数 * cfs_period_us
	cpuCfsPeriod = "cpu.cfs_period_us"
	// 表示Cgroup可以使用的cpu的带宽，单位为微秒。cfs_quota_us为-1，表示使用的CPU不受cgroup限制。cfs_quota_us的最小值为1ms(1000)，最大值为1s。
	cpuCfsQuota  = "cpu.cfs_quota_us"
	cpuRtPeriod  = "cpu.rt_period_us"
	cpuRtRuntime = "cpu.rt_runtime_us"
	//通过cfs_period_us和cfs_quota_us可以以绝对比例限制cgroup的cpu使用，即cfs_quota_us/cfs_period_us 等于进程可以利用的cpu cores，不能超过这个数值。
	cpuShares = "cpu.shares"
)

func SetCpuCGroup(cpuPath string, pid int) error {
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
