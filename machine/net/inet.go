package net

type INetCard interface {
	GetName() string  // 获取网卡名称
	GetMtu() int      // 获取网卡数据包大小
	GetMac() string   // 获取网卡mac地址
	GetIPV4() string  // 获取网卡IP
	GetState() string // 获取网卡状态
	ToString() string // 获取json信息
}

type INet interface {
	GetCard(name string) INetCard // 获取单个网卡
	GetCardList() []string        // 获取网卡列表
	ToString() string             // 获取网卡json数据
}
