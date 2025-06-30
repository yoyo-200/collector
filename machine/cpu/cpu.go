package cpu

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/yoyo-200/collector/machine/base"
	"github.com/yoyo-200/collector/machine/formatsize"
	"github.com/yoyo-200/collector/utils"

	cpus "github.com/shirou/gopsutil/cpu"
)

// cpu信息
type cpu struct {
	Name         string    `json:"name"`                                                   // cpu 名称
	Arch         string    `json:"arch"`                                                   // cpu 架构
	Cores        int64     `json:"core"`                                                   // cpu 逻辑核心数
	OverallUsage string    `json:"overall_usage"`                                          // 总体CPU使用率百分比
	PerCPUUsage  *[]string `json:"per_cpu_usage,omitempty" yaml:"per_cpu_usage,omitempty"` // 每个CPU核心的使用率
}

// 配置结构体
type cpuConfig struct {
	percpu   bool          // 是否获取每个核心的信息
	interval time.Duration // 采样时间间隔
}

// 配置选项：控制是否获取每个cpu得值,false则是单个cpu值
type CPUOption func(*cpuConfig)

// 默认配置
func defaultCpuConfig() *cpuConfig {
	return &cpuConfig{
		percpu:   false, // 默认不包含所有cpu值,百分比显示平均值
		interval: 0,     // 默认采集1秒时间间隔
	}
}

// 选项函数：是否包含所有CPU值 默认是false不包含所有
func WithAllCpu(include bool) CPUOption {
	return func(c *cpuConfig) {
		c.percpu = include
	}
}

// NewCPU 创建CPU信息对象
// 默认配置：只获取总体使用率，采样间隔1秒
func NewCPU(options ...CPUOption) ICpu {
	// 默认配置
	config := defaultCpuConfig()
	// 选项
	for _, option := range options {
		option(config)
	}

	// 获取总核心数，默认是逻辑核心数
	cores, err := cpus.Counts(true)
	if err != nil {
		utils.DefaultLogger.Error(err)
	}

	// 获取CPU使用率
	percent, err := cpus.Percent(config.interval, config.percpu)
	if err != nil {
		utils.DefaultLogger.Error(err)
	}

	// 获取CPU信息
	stats, err := cpus.Info()
	if err != nil {
		utils.DefaultLogger.Error(err)
	}

	// 构造结果
	info := &cpu{
		Name:  stats[0].ModelName,
		Arch:  base.NewBaseInfo().GetArch(),
		Cores: int64(cores),
	}

	if config.percpu {
		// 返回所有核心的信息
		perCPU := make([]string, 0, len(percent))
		for _, p := range percent {
			perCPU = append(perCPU, formatsize.FormatPercent(p))
		}
		info.PerCPUUsage = &perCPU
		if len(percent) > 0 {
			// 计算总体使用率（所有核心平均值）
			var total float64
			for _, p := range percent {
				total += p
			}
			info.OverallUsage = fmt.Sprintf("%.2f%%", total/float64(len(percent)))
		}
	} else {
		// 只返回总体使用率
		if len(percent) > 0 {
			info.OverallUsage = fmt.Sprintf("%.2f%%", percent[0])
		}
	}

	return info

}

func (c *cpu) GetCpuName() string { return c.Name }  // 获取cpu名称
func (c *cpu) GetCpuArch() string { return c.Arch }  // 获取cpu 架构
func (c *cpu) GetCpuCores() int64 { return c.Cores } // 获取cpu 逻辑核心数

// 获取cpu 物理核心数
func (c *cpu) GetCpuPhysicalCores() int64 {
	counts, err := cpus.Counts(false)
	if err != nil {
		utils.DefaultLogger.Error(err)
	}
	return int64(counts)
}

// 获取cpu使用率(该方法是单独想获取某个时间段的cpu使用率)
func (c *cpu) GetTimeToCpuPercent(interval time.Duration, isAllCpu bool) []string {
	percents, err := cpus.Percent(interval, isAllCpu)
	if err != nil {
		utils.DefaultLogger.Error(err)
	}

	p := []string{}

	for _, percent := range percents {
		per := formatsize.FormatPercent(percent)
		p = append(p, per)
	}

	return p
}

// 获取cpu json信息
func (c *cpu) ToString() string {
	cpuJson, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		utils.DefaultLogger.Error(err)
	}
	return string(cpuJson)
}
