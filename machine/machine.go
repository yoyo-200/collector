package machine

import (
	"encoding/json"

	"github.com/yoyo-200/collector/machine/base"
	"github.com/yoyo-200/collector/machine/cpu"
	"github.com/yoyo-200/collector/machine/disk"
	"github.com/yoyo-200/collector/machine/memory"
	"github.com/yoyo-200/collector/machine/net"

	"gopkg.in/yaml.v3"
)

type Machine struct {
	BaseInfo base.IBaseInfo `json:"base" yaml:"base"`       // 基础信息
	Cpu      cpu.ICpu       `json:"cpu" yaml:"cpu"`         // cpu信息
	Disk     disk.IDisk     `json:"disk" yaml:"disk"`       // disk信息
	Memory   memory.IMemory `json:"memory" yaml:"memory"`   // memory信息
	Network  net.INet       `json:"network" yaml:"network"` // net信息
	config   *MachineConfig // 保存配置选项
}

// 配置选项：控制
type MachineOption func(*MachineConfig)

// 配置结构体
type MachineConfig struct {
	isAllCpuInfo  bool     // 是否包涵所有cpu信息
	isAllDiskInfo bool     // 是否包涵有disk信息
	isShipFsType  []string // 是否有跳过的文件类型
}

// 默认配置
func DefaultMachineConfig() *MachineConfig {
	return &MachineConfig{
		isAllCpuInfo:  false, // 默认是false 不包含所有cpu信息
		isAllDiskInfo: true,  // 默认是true 包含所有disk信息(主要是包含disk中虚拟机磁盘信息)
		isShipFsType: []string{
			"pstore", "pstorefs", "sysfs", "cgroup", "cgroup2", "autofs",
			"mqueue", "devpts", "hugetlbfs", "configfs",
			"debugfs", "tracefs", "securityfs", "efivarfs",
			"fusectl", "binfmt_misc", "rpc_pipefs",
			"nsfs", "selinuxfs", "fuse.gvfsd-fuse", "proc", ""}, // 默认跳过得虚拟文件系统类型
	}
}

// 是否包含所有CPU值,默认false,不包含所有cpu信息
func WithAllCpu(include bool) MachineOption {
	return func(ma *MachineConfig) {
		ma.isAllCpuInfo = include
	}
}

// 是否包含所有磁盘,默认是true.包含虚拟磁盘
func WithAllDisk(include bool) MachineOption {
	return func(ma *MachineConfig) {
		ma.isAllDiskInfo = include
	}
}

// 是否跳过disk中哪些虚拟磁盘文件系统类型
// 默认: "pstore", "pstorefs", "sysfs", "cgroup", "cgroup2", "autofs","mqueue", "devpts", "hugetlbfs",
// "configfs","debugfs", "tracefs", "securityfs", "efivarfs","fusectl", "binfmt_misc", "rpc_pipefs","nsfs", "selinuxfs", "fuse.gvfsd-fuse", "proc",
func WithSkipFstype(include []string) MachineOption {
	return func(ma *MachineConfig) {
		ma.isShipFsType = include
	}
}

func NewMachine(options ...MachineOption) IMachine {

	// 默认配置,新增选择进行覆盖
	config := DefaultMachineConfig()
	// 选项
	for _, option := range options {
		option(config)
	}
	return &Machine{
		BaseInfo: base.NewBaseInfo(),
		Cpu:      cpu.NewCPU(),
		Disk:     disk.NewDisk(disk.WithVirtual(config.isAllDiskInfo), disk.WithSkipFstyp(config.isShipFsType)),
		Memory:   memory.NewMemory(),
		Network:  net.NewNet(),
		config:   config,
	}
}

// 获取基础信息
func (m *Machine) BASE() base.IBaseInfo { return m.BaseInfo }

// 获取cpu信息
func (m *Machine) CPU() cpu.ICpu {
	// 动态加载cpu信息，当为true 获取所有cpu信息
	if m.config.isAllCpuInfo {
		m.Cpu = cpu.NewCPU(cpu.WithAllCpu(m.config.isAllCpuInfo))
	}
	// 返回默认配置，不包含所有cpu
	return m.Cpu
}

// 获取磁盘信息
func (m *Machine) DISK() disk.IDisk {
	// 动态加载disk，当为false 仅获取物理磁盘信息
	if !m.config.isAllDiskInfo {
		m.Disk = disk.NewDisk(disk.WithVirtual(false))
	}
	// else {
	// 	// 返回默认配置，包含所有disk信息(主要是包含disk中虚拟机磁盘信息)
	// 	m.Disk = disk.NewDisk(disk.WithVirtual(m.config.isAllDiskInfo), disk.WithSkipFstyp(m.config.isShipFsType))
	// }
	return m.Disk
}

// 获取内存信息
func (m *Machine) MEMORY() memory.IMemory { return m.Memory }

// 获取网卡信息
func (m *Machine) NETWORK() net.INet { return m.Network }

// 获取json数据
func (m *Machine) ToJsonString() string {
	data, _ := json.MarshalIndent(m, "", " ")
	// utils.DefaultLogger.Error(err)
	return string(data)
}

// 获取yaml数据
func (m *Machine) ToYamlString() string {
	datayaml, _ := yaml.Marshal(&m)
	// utils.DefaultLogger.Error(err)
	return string(datayaml)
}
