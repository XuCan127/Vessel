package cgroup

import (
	"Vessel/src/server/module/cgroup/subSystems"
	"path"
)

const (
	CGroupPath = "/sys/fs/cgroup"
)

func SetCGroup(pid int, containerName string, res *subSystems.ResourceConfig) error {

	var (
		cpuPath    = path.Join(CGroupPath, "cpu", containerName)
		memoryPath = path.Join(CGroupPath, "memory", containerName)
	)
	// 设置cpu
	if err := subSystems.SetCpuCGroup(cpuPath, pid, res); err != nil {
		panic(err)
	}
	// 设置内存
	if err := subSystems.SetMemoryCGroup(memoryPath, pid, res); err != nil {
		panic(err)
	}
	return nil
}

func CleanCGroup(containerName string) error {

	var (
		cpuPath    = path.Join(CGroupPath, "cpu", containerName)
		memoryPath = path.Join(CGroupPath, "memory", containerName)
	)
	if err := subSystems.CleanCpuCGroup(cpuPath); err != nil {
		panic(err)
	}
	if err := subSystems.CleanMemoryCGroup(memoryPath); err != nil {
		panic(err)
	}
	return nil

}
