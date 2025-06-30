package disk

type iOneDisk interface {
	GetPath() string        // 获取单个磁盘路径
	GetTotal() string       // 获取单个磁盘总空间
	GetFree() string        // 获取单个磁盘空闲空间
	GetReadWrite() string   // 获取单个磁盘读写权限
	GetUsed() string        // 获取单个磁盘使用空间
	GetUsedPercent() string // 获取单个磁盘使用占比
	GetDevice() string      // 获取单个磁盘文件系统驱动
	GetFstype() string      // 获取单个磁盘文件系统
	GetInodesTotal() string // 获取单个磁盘inode总空间
	GetInodesUsed() string  // 获取单个磁盘inode使用空间
	GetInodesFree() string  // 获取单个磁盘inode空闲空间
	GetMountPoint() string  // 获取单个磁盘挂载点
}

type IDisk interface {
	GetDiskPath(diskName string) string // 获取磁盘路径
	GetDiskList() []string              // 获取磁盘列表
	ToString() string                   // 获取json数据
}
