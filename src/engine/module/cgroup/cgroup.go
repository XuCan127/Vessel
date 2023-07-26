package cgroup

import (
	"github.com/XuCan127/Vessel/src/engine/module/cgroup/subSystems"
	"path"
)

const (
	CGroupPath = "/sys/fs/cgroup"
)

func SetCGroup(pid int, containerName string) error {

	var (
		cpuPath    = path.Join(CGroupPath, "cpu", containerName)
		memoryPath = path.Join(CGroupPath, "memory", containerName)
	)
	// 设置cpu
	if err := subsystem.SetCpuCGroup(cpuPath, pid); err != nil {
		panic(err)
	}
	// 设置内存
	if err := subsystem.SetMemoryCGroup(memoryPath, pid); err != nil {
		panic(err)
	}
	return nil
}

func CleanCGroup(containerName string) error {

	var (
		cpuPath    = path.Join(CGroupPath, "cpu", containerName)
		memoryPath = path.Join(CGroupPath, "memory", containerName)
	)
	if err := subsystem.CleanCpuCGroup(cpuPath); err != nil {
		panic(err)
	}
	if err := subsystem.CleanMemoryCGroup(memoryPath); err != nil {
		panic(err)
	}
	return nil

}
