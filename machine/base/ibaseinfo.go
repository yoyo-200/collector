package base

import "time"

type IBaseInfo interface {
	GetArch() string          // 获取架构信息
	GetPlatform() string      // 获取系统平台
	GetOs() string            // 获取os
	GetKernel() string        // 获取内核版本
	GetIPv4() string          // 获取ip地址
	GetSystemType() string    // 获取系统类型
	GetHostname() string      // 获取主机名
	GetUUID() string          // 获取系统UUID
	GetVirtual() string       // 获取系统是虚拟机还是物理机
	GetBootTime() string      // 获取boot time
	GetUpdateTime() time.Time // 获取更新时间
	ToString() string         // 获取json信息
}
