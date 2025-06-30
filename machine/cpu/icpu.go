package cpu

import "time"

type ICpu interface {
	GetCpuName() string                                            // 获取cpu名称
	GetCpuArch() string                                            // 获取cpu 架构
	GetCpuCores() int64                                            // 获取cpu 核心数
	GetCpuPhysicalCores() int64                                    // 获取cpu 物理核心数
	GetTimeToCpuPercent(times time.Duration, allCPU bool) []string // 获取某时间段cpu使用率
	ToString() string                                              // 获取json数据
}
