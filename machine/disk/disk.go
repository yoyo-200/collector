package disk

import (
	"encoding/json"
	"runtime"
	"strings"

	"github.com/yoyo-200/collector/machine/formatsize"
	"github.com/yoyo-200/collector/utils"

	"github.com/shirou/gopsutil/disk"
)

// 单个磁盘信息
type oneDisk struct {
	Path        string `json:"path"`                                               // 磁盘路径
	Total       string `json:"total"`                                              // 磁盘总量
	Free        string `json:"free"`                                               // 磁盘空闲量
	Permission  string `json:"permission"`                                         // 磁盘读写权限
	Used        string `json:"used"`                                               // 磁盘使用量
	UsedPercent string `json:"usedpercent"`                                        // 使用率
	Device      string `json:"device"`                                             // 磁盘驱动
	Fstype      string `json:"fstype,omitempty" yaml:"fstype,omitempty"`           // 磁盘文件系统
	InodesTotal string `json:"inodesTotal,omitempty" yaml:"inodesTotal,omitempty"` // 磁盘 inode 数
	InodesUsed  string `json:"inodesUsed,omitempty" yaml:"inodesUsed,omitempty"`   // 磁盘已用 inode 数
	InodesFree  string `json:"inodesFree,omitempty" yaml:"inodesFree,omitempty"`   // 磁盘空闲 inode 数
	MountPoint  string `json:"mountpoint"`                                         // 磁盘挂载点
}

func (od *oneDisk) GetPath() string        { return od.Path }        // 获取单个磁盘路径
func (od *oneDisk) GetTotal() string       { return od.Total }       // 获取单个磁盘总空间
func (od *oneDisk) GetFree() string        { return od.Free }        // 获取单个磁盘空闲空间
func (od *oneDisk) GetReadWrite() string   { return od.Permission }  // 获取单个磁盘读写权限
func (od *oneDisk) GetUsed() string        { return od.Used }        // 获取单个磁盘使用空间
func (od *oneDisk) GetUsedPercent() string { return od.UsedPercent } // 获取单个磁盘使用占比
func (od *oneDisk) GetDevice() string      { return od.Device }      // 获取单个磁盘文件系统驱动
func (od *oneDisk) GetFstype() string      { return od.Fstype }      // 获取单个磁盘文件系统
func (od *oneDisk) GetInodesTotal() string { return od.InodesTotal } // 获取单个磁盘inode总空间
func (od *oneDisk) GetInodesUsed() string  { return od.InodesUsed }  // 获取单个磁盘inode使用空间
func (od *oneDisk) GetInodesFree() string  { return od.InodesFree }  // 获取单个磁盘inode空闲空间
func (od *oneDisk) GetMountPoint() string  { return od.MountPoint }  // 获取单个磁盘挂载点

// 磁盘信息 格式： /dev/sda1: 20G
type Disk map[string]*oneDisk

// 配置选项：控制是否包含虚拟机磁盘（如虚拟文件系统）
type DiskOption func(*diskConfig)

// 配置结构体
type diskConfig struct {
	includeVirtual bool     // 是否包含虚拟文件系统（如 tmpfs）默认是包含虚拟文件系统
	skipFsType     []string // 跳过特殊文件类型得清单
}

// 默认配置
func defaultDiskConfig() *diskConfig {
	return &diskConfig{
		includeVirtual: true, // 默认包含虚拟磁盘
		skipFsType:     nil,  // 默认跳过类型,默认是空
	}
}

// NewDisk 创建DISK信息对象
func NewDisk(opt ...DiskOption) IDisk {
	// 默认配置
	config := defaultDiskConfig()
	// 选项
	for _, option := range opt {
		option(config)
	}

	// 构造对象
	diskInfo := make(Disk)
	partitions, err := disk.Partitions(config.includeVirtual)
	if err != nil {
		utils.DefaultLogger.Error(err)
	}
	// 检测当前操作系统
	isWindows := runtime.GOOS == "windows"
	for _, partition := range partitions {
		// 检查是否需要跳过当前文件系统
		if config.skipFsType != nil {
			skip := false
			for _, fsType := range config.skipFsType {
				// 增强：处理空文件系统类型
				if partition.Device == "none" || partition.Device == "bpf" {
					skip = true
					break
				}
				// 使用大小写不敏感的比较
				if strings.EqualFold(partition.Fstype, fsType) {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
		}
		// 获取磁盘使用信息（包含 inode）
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}
		// 创建完整的磁盘信息
		diskInfo[partition.Mountpoint] = &oneDisk{
			Path:        usage.Path,
			Total:       formatsize.FormatSize(usage.Total),
			Free:        formatsize.FormatSize(usage.Free),
			Permission:  partition.Opts[:2],
			Used:        formatsize.FormatSize(usage.Used),
			UsedPercent: formatsize.FormatPercent(usage.UsedPercent),
			Device:      partition.Device,
			Fstype:      usage.Fstype,
			MountPoint:  partition.Mountpoint,
		}

		// 如果不是 Windows 平台，添加 inode 信息
		if !isWindows {
			diskInfo[partition.Mountpoint].InodesTotal = formatsize.FormatSize(usage.InodesTotal)
			diskInfo[partition.Mountpoint].InodesUsed = formatsize.FormatSize(usage.InodesUsed)
			diskInfo[partition.Mountpoint].InodesFree = formatsize.FormatSize(usage.InodesFree)
		}
	}
	return &diskInfo
}

// 跳过不需要显示的特殊文件系统
func WithSkipFstyp(skipfstype []string) DiskOption {
	return func(c *diskConfig) {
		c.skipFsType = skipfstype
	}
}

// 默认是true.包含虚拟磁盘
func WithVirtual(include bool) DiskOption {
	return func(c *diskConfig) {
		c.includeVirtual = include
	}
}

// 获取磁盘路径
func (d *Disk) GetDiskPath(diskName string) string {
	if d == nil || *d == nil { // 双重 nil 检查
		return ""
	}
	if disk, ok := (*d)[diskName]; ok && disk != nil {
		return disk.GetPath()
	}
	return ""

}

func (d *Disk) GetDiskList() []string {
	disklist := []string{}

	for _, i := range *d {
		disklist = append(disklist, i.Device)
	}
	return disklist
} // 获取磁盘列表

// 获取json数据
func (d *Disk) ToString() string {
	diskJson, err := json.MarshalIndent(d, "", " ")
	if err != nil {
		utils.DefaultLogger.Error(err)
	}
	return string(diskJson)
}
