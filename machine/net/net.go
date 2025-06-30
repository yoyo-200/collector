package net

import (
	"encoding/json"
	"net"
	"runtime"

	"gitee.com/yolzj/collector/utils"
)

// 单网卡信息
type networkCard struct {
	Name  string `json:"name"`                               // 网卡名称
	Mtu   int    `json:"mtu"`                                // 网卡接收/发送最大数据包大小
	Mac   string `json:"mac,omitempty" yaml:"mac,omitempty"` // mac地址
	IP    string `json:"ip"`                                 // 默认ipv4地址,没有ipv4则显示ipv6地址
	State string `json:"state"`                              // 网卡状态
	// 未实现
	// Speed               string   `json:"speed"`                 // 网卡速率
	// LinkDetected        string   `json:"link_detected"`         // 链路检测状态
	// Duplex              string   `json:"duplex"`                // 双工模式
	// SupportedLinkModes  []string `json:"supported_link_modes"`  // 支持的链路模式
	// AdvertisedLinkModes []string `json:"advertised_link_modes"` // 宣告的链路模式
}

func (nc *networkCard) GetName() string  { return nc.Name }  // 获取网卡名称
func (nc *networkCard) GetMtu() int      { return nc.Mtu }   // 获取网卡数据包大小
func (nc *networkCard) GetMac() string   { return nc.Mac }   // 获取网卡mac地址
func (nc *networkCard) GetIPV4() string  { return nc.IP }    // 获取网卡IP
func (nc *networkCard) GetState() string { return nc.State } // 获取网卡状态

// 网卡 json 信息
func (nc *networkCard) ToString() string {
	b, err := json.MarshalIndent(nc, "", " ")
	if err != nil {
		utils.DefaultLogger.Error(err)
	}
	return string(b)
}

// 网卡合集
type Net map[string]*networkCard // 网卡信息(所有网卡)

// 网卡合集句柄
func NewNet() INet {
	netcards := make(Net)
	n, err := net.Interfaces()
	if err != nil {
		utils.DefaultLogger.Error(err)
	}

	for _, i := range n {
		ip := ""
		addrs, err := i.Addrs()
		if err != nil {
			utils.DefaultLogger.Error(err)
		}
		if len(addrs) > 0 {
			if runtime.GOOS == "windows" {
				ip = addrs[1].(*net.IPNet).IP.String()
			} else {
				ip = addrs[0].(*net.IPNet).IP.String()
			}

		}

		// 构造数据
		networkinfo := networkCard{
			Name:  i.Name,
			Mtu:   i.MTU,
			Mac:   i.HardwareAddr.String(),
			IP:    ip,
			State: i.Flags.String(),
		}

		netcards[i.Name] = &networkinfo
	}
	return &netcards
}

// 获取单个网卡
func (n *Net) GetCard(name string) INetCard {
	for _, i := range *n {
		if i.Name == name {
			return i
		}
	}
	return nil
}

// 获取网卡列表
func (n *Net) GetCardList() []string {
	networkcards := []string{}
	for _, i := range *n {
		networkcards = append(networkcards, i.Name)
	}
	return networkcards
}

func (n *Net) ToString() string {
	data, err := json.MarshalIndent(n, "", " ")
	utils.DefaultLogger.Error(err)
	return string(data)
} // 获取网卡json数据
