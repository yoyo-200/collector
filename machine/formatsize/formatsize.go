package formatsize

import (
	"fmt"
	"time"
)

// func FormatSize[T constraints.Integer](size T) string {

// 	sizeInBytes := float64(size)

// 	if sizeInBytes < 1024 {
// 		return fmt.Sprintf("%.2f B", sizeInBytes)
// 	}
// 	if sizeInBytes < 1024*1024 {
// 		return fmt.Sprintf("%.2f KB", sizeInBytes/1024)
// 	}
// 	if sizeInBytes < 1024*1024*1024 {
// 		return fmt.Sprintf("%.2f MB", sizeInBytes/(1024*1024))
// 	} else {
// 		return fmt.Sprintf("%.2f GB", sizeInBytes/(1024*1024*1024))
// 	}
// }

func FormatSize(size uint64) string {
	sizeInBytes := float64(size)

	if sizeInBytes < 1024 {
		return fmt.Sprintf("%.2f B", sizeInBytes)
	}
	if sizeInBytes < 1024*1024 {
		return fmt.Sprintf("%.2f KB", sizeInBytes/1024)
	}
	if sizeInBytes < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", sizeInBytes/(1024*1024))
	} else {
		return fmt.Sprintf("%.2f GB", sizeInBytes/(1024*1024*1024))
	}
}

func FormatPercent(size float64) string {
	return fmt.Sprintf("%.2f%%", size)
}

// 新曾格式化uint64 >> 时间戳
func FormatTime(size uint64) string {
	// 创建中国标准时间时区 (UTC+8)
	cstZone := time.FixedZone("CST", 8*60*60)
	// 转换为 time.Time 对象
	t := time.Unix(int64(size), 0).In(cstZone)

	// 格式化为标准时间字符串
	return t.Format("2006-01-02 15:04:05 MST")
}
