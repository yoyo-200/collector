package template

import (
	"gitee.com/yolzj/collector/machine"
	"gitee.com/yolzj/collector/machine/cpu"
	"gitee.com/yolzj/collector/machine/disk"
	"gitee.com/yolzj/collector/machine/memory"
	"gitee.com/yolzj/collector/machine/net"
)

type Template struct {
	//添加自定义模版自段，此处仅为示例直接继承machine.IMachine
	machine machine.IMachine
	// 模板特有字段
	UUID         string `json:"uuid"`          // UUID
	Hostname     string `json:"hostname"`      // 主机名
	IP           string `json:"ip"`            // ip
	Type         string `json:"type"`          // 类型
	Environment  string `json:"environment"`   // 环境
	OSType       string `json:"os_type"`       // OS类型
	Version      string `json:"version"`       // 系统版本
	SystemType   string `json:"system_type"`   // 系统类型
	HostPassword string `json:"host_password"` // 密码 (主机相关)
	BIOSTime     string `json:"bios_time"`     // bios时间
	UpdateTime   string `json:"update_time"`   // 更新时间
}

// NewOne 构造一个模版“one”实例
func NewOne(options ...machine.MachineOption) ITemplate {

	newMachine := machine.NewMachine(options...)
	return &Template{
		machine:     newMachine,
		UUID:        newMachine.BASE().GetUUID(),
		Hostname:    newMachine.BASE().GetHostname(),
		IP:          newMachine.BASE().GetIPv4(),
		Type:        newMachine.BASE().GetArch(),
		Environment: newMachine.BASE().GetVirtual(),
		OSType:      newMachine.BASE().GetPlatform(),
		Version:     newMachine.BASE().GetOs(),
		SystemType:  newMachine.BASE().GetSystemType(),
		BIOSTime:    newMachine.BASE().GetBootTime(),
		UpdateTime:  newMachine.BASE().GetUpdateTime().String(),
	}
}

/*
动态获取cpu memory disk network信息
*/
func (t *Template) GetCpu() cpu.ICpu          { return t.machine.CPU() }     // 模板cpu
func (t *Template) GetMemory() memory.IMemory { return t.machine.MEMORY() }  // 模板内存
func (t *Template) GetDisk() disk.IDisk       { return t.machine.DISK() }    // 模板磁盘
func (t *Template) GetNetWork() net.INet      { return t.machine.NETWORK() } // 模板网卡

func (t *Template) GetUUID() string         { return t.UUID }         // 模板UUID
func (t *Template) GetHostname() string     { return t.Hostname }     // 模板主机名
func (t *Template) GetIP() string           { return t.IP }           // 模板ip
func (t *Template) GetType() string         { return t.Type }         // 模板类型
func (t *Template) GetEnvironment() string  { return t.Environment }  // 模板环境
func (t *Template) GetOSType() string       { return t.OSType }       // 模板OS类型
func (t *Template) GetVersion() string      { return t.Version }      // 模板系统版本
func (t *Template) GetSystemType() string   { return t.SystemType }   // 模板系统类型
func (t *Template) GetHostPassword() string { return t.HostPassword } // 模板密码 (主机相关)
func (t *Template) GetBIOSTime() string     { return t.BIOSTime }     // 模板bios时间
func (t *Template) GetUpdateTime() string   { return t.UpdateTime }   // 模板更新时间
