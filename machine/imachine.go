package machine

import (
	"gitee.com/yolzj/collector/machine/base"
	"gitee.com/yolzj/collector/machine/cpu"
	"gitee.com/yolzj/collector/machine/disk"
	"gitee.com/yolzj/collector/machine/memory"
	"gitee.com/yolzj/collector/machine/net"
)

type IMachine interface {
	BASE() base.IBaseInfo   // 获取机器基础信息
	CPU() cpu.ICpu          // 获取机器cpu信息
	DISK() disk.IDisk       // 获取机器磁盘信息
	MEMORY() memory.IMemory // 获取机器内存信息
	NETWORK() net.INet      // 获取机器网卡信息
	ToJsonString() string   // 获取json数据
	ToYamlString() string   // 获取yaml数据
}
