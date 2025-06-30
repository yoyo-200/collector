package main

import (
	"fmt"

	"github.com/yoyo-200/collector/machine/template"
)

func main() {
	// 创建模版“one”实例
	t := template.NewOne()
	cpu := t.GetCpu()
	name := cpu.GetCpuName()
	fmt.Printf("cpu: %v\n", name)

	disk := t.GetDisk()
	disk.ToString()
}
