package base

import (
	"encoding/json"
	"log"
	"runtime"
	"time"

	"github.com/elastic/go-sysinfo"
	"github.com/shirou/gopsutil/host"
	machine "github.com/super-l/machine-code"
	"github.com/yoyo-200/collector/machine/formatsize"
	"github.com/yoyo-200/collector/utils"
)

// 系统基础信息
type baseInfo struct {
	UUID          string    `json:"uuid"`                                                     // uuid
	Platform      string    `json:"platform"`                                                 // 属于那个家族例如: rocky debian
	Os            string    `json:"os"`                                                       // 发行版
	Kernel        string    `json:"kernel"`                                                   // 内核
	Arch          string    `json:"arch"`                                                     // 架构
	Hostname      string    `json:"hostname"`                                                 // 主机名
	IPv4          string    `json:"ipv4"`                                                     // ip
	SystemType    string    `json:"systemtype"`                                               // 系统类型 例如: windows linux
	VirtualSystem string    `json:"virtual_system,omitempty" yaml:"virtual_system,omitempty"` // 系统里的虚拟机化平台
	VirtualRole   string    `json:"virtual_role,omitempty" yaml:"virtual_role,omitempty"`     // 是宿主机还是虚拟机
	BootTime      string    `json:"boot_time" yaml:"boot_time"`                               // boot 时间
	UpdateTime    time.Time `json:"update_time"`                                              // 更新时间
}

// cpu 模块内以这个私有变量管理
var basesInfo = baseInfoLoad()

// 返回gopsutils host.InfoStat对象
func baseInfoLoad() *host.InfoStat {
	baseInfo, err := host.Info()
	if err != nil {
		utils.DefaultLogger.Error(err)
	}
	return baseInfo
}

// 获取当前主机信息
func getSystemVersion() string {

	host, err := sysinfo.Host()
	if err != nil {
		panic(err)
	}
	os := host.Info()
	// os_version :=
	return os.OS.Name + "" + os.OS.Version
}

// 获取ip
func getIPv4Addr() string {
	// 获取ip
	addr, err := machine.GetIpAddr()
	if err != nil {
		log.Fatal(err)
	}
	return addr
}

// NewBaseInfo 系统基础信息句柄
func NewBaseInfo() IBaseInfo {

	return &baseInfo{
		UUID:          basesInfo.HostID,
		Platform:      basesInfo.Platform,
		Os:            getSystemVersion(),
		Kernel:        basesInfo.KernelVersion,
		Arch:          basesInfo.KernelArch,
		Hostname:      basesInfo.Hostname,
		IPv4:          getIPv4Addr(),
		SystemType:    runtime.GOOS,
		VirtualSystem: basesInfo.VirtualizationSystem,
		VirtualRole:   basesInfo.VirtualizationRole,
		BootTime:      formatsize.FormatTime(basesInfo.BootTime),
		UpdateTime:    time.Now(),
	}
}

func (b *baseInfo) GetArch() string          { return b.Arch }       // 获取架构信息
func (b *baseInfo) GetPlatform() string      { return b.Platform }   // 获取系统平台
func (b *baseInfo) GetOs() string            { return b.Os }         // 获取os
func (b *baseInfo) GetKernel() string        { return b.Kernel }     // 获取内核版本
func (b *baseInfo) GetIPv4() string          { return b.IPv4 }       // 获取ip地址
func (b *baseInfo) GetSystemType() string    { return b.SystemType } // 获取系统类型
func (b *baseInfo) GetHostname() string      { return b.Hostname }   // 获取主机名
func (b *baseInfo) GetUUID() string          { return b.UUID }       // 获取系统UUID
func (b *baseInfo) GetBootTime() string      { return b.BootTime }   // 获取boot time
func (b *baseInfo) GetUpdateTime() time.Time { return b.UpdateTime } // 获取更新时间

// 获取系统是虚拟机还是物理机
func (b *baseInfo) GetVirtual() string {
	if b.VirtualRole == "host" {
		return "Physical"
	}
	return "Virtual"
}

// 获取json信息
func (b *baseInfo) ToString() string {
	data, err := json.MarshalIndent(b, "", " ")
	if err != nil {
		utils.DefaultLogger.Error(err)
	}
	return string(data)
}
