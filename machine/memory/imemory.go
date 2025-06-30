package memory

type oneTypeMemory interface {
	GetTotal() string       // 获取内存总量
	GetUsed() string        // 获取内存使用量
	GetFree() string        // 获取内存空闲量
	GetUsedPercent() string // 获取内存使用百分比
	ToString() string       // 获取json数据
}

type IMemory interface {
	GetVirtMemory() oneTypeMemory // 获取virt对象
	GetSwapMemory() oneTypeMemory // 获取swap对象
	ToString() string             // 获取json数据
}
