package template

import (
	"github.com/yoyo-200/collector/machine/cpu"
	"github.com/yoyo-200/collector/machine/disk"
	"github.com/yoyo-200/collector/machine/memory"
	"github.com/yoyo-200/collector/machine/net"
)

type ITemplate interface {
	GetUUID() string           // 模板UUID
	GetHostname() string       // 模板主机名
	GetIP() string             // 模板ip
	GetType() string           // 模板类型
	GetEnvironment() string    // 模板环境
	GetOSType() string         // 模板OS类型
	GetVersion() string        // 模板系统版本
	GetSystemType() string     // 模板系统类型
	GetHostPassword() string   // 模板密码 (主机相关)
	GetCpu() cpu.ICpu          // 模板cpu
	GetMemory() memory.IMemory // 模板内存
	GetDisk() disk.IDisk       // 模板磁盘
	GetNetWork() net.INet      // 模板网卡
	GetBIOSTime() string       // 模板bios时间
	GetUpdateTime() string     // 模板更新时间
}
