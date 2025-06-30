package machine

import (
	"github.com/yoyo-200/collector/machine/base"
	"github.com/yoyo-200/collector/machine/cpu"
	"github.com/yoyo-200/collector/machine/disk"
	"github.com/yoyo-200/collector/machine/memory"
	"github.com/yoyo-200/collector/machine/net"
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
