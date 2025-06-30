package memory

import (
	"encoding/json"
	"fmt"

	"gitee.com/yolzj/collector/machine/formatsize"
	"gitee.com/yolzj/collector/utils"

	"github.com/shirou/gopsutil/mem"
)

// 虚拟内存信息
type virtualMemory struct {
	Total       string `json:"total"`
	Used        string `json:"used"`
	Free        string `json:"free"`
	UsedPercent string `json:"usedpercent"`
}

func (vm *virtualMemory) GetTotal() string       { return vm.Total }       // 获取内存总量
func (vm *virtualMemory) GetUsed() string        { return vm.Used }        // 获取内存使用量
func (vm *virtualMemory) GetFree() string        { return vm.Free }        // 获取内存空闲量
func (vm *virtualMemory) GetUsedPercent() string { return vm.UsedPercent } // 获取内存使用百分比

// 获取json数据
func (vm *virtualMemory) ToString() string {
	jsonString, _ := json.Marshal(vm)
	return string(jsonString)
}

// 交换分区内存信息
type swapMemory struct {
	Total       string `json:"total"`
	Used        string `json:"used"`
	Free        string `json:"free"`
	UsedPercent string `json:"usedpercent"`
}

func (sm *swapMemory) GetTotal() string       { return sm.Total }       // 获取内存总量
func (sm *swapMemory) GetUsed() string        { return sm.Used }        // 获取内存使用量
func (sm *swapMemory) GetFree() string        { return sm.Free }        // 获取内存空闲量
func (sm *swapMemory) GetUsedPercent() string { return sm.UsedPercent } // 获取内存使用百分比

// 获取json数据
func (sm *swapMemory) ToString() string {
	jsonString, _ := json.Marshal(sm)
	return string(jsonString)
}

// 内存信息
type memory struct {
	Virt oneTypeMemory `json:"virtual_memory" yaml:"virtual_memory"` // 虚拟内存信息
	Swap oneTypeMemory `json:"swap_memory" yaml:"swap_memory"`       // 交换分区内存信息
}

// memory 句柄
func NewMemory() IMemory {
	// 获取虚拟内存信息
	virtmemory, err := mem.VirtualMemory()
	if err != nil {
		utils.DefaultLogger.Error(err)
	}

	// 获取交换区内存信息
	swapmemory, err := mem.SwapMemory()
	if err != nil {
		utils.DefaultLogger.Error(err)
	}
	return &memory{
		Virt: &virtualMemory{
			Total:       formatsize.FormatSize(virtmemory.Total),
			Used:        formatsize.FormatSize(virtmemory.Used),
			Free:        formatsize.FormatSize(virtmemory.Free),
			UsedPercent: fmt.Sprintf("%.2f%%", virtmemory.UsedPercent),
		},
		Swap: &swapMemory{
			Total:       formatsize.FormatSize(swapmemory.Total),
			Used:        formatsize.FormatSize(swapmemory.Used),
			Free:        formatsize.FormatSize(swapmemory.Free),
			UsedPercent: fmt.Sprintf("%.2f%%", swapmemory.UsedPercent),
		},
	}
}

func (m *memory) GetVirtMemory() oneTypeMemory { return m.Virt } // 获取虚拟内存信息
func (m *memory) GetSwapMemory() oneTypeMemory { return m.Swap } // 获取swap内存信息

// 获取memory json信息
func (m *memory) ToString() string {
	data, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		utils.DefaultLogger.Error(err)
	}
	return string(data)
}
