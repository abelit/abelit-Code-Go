package main

import (
	"fmt"
	"runtime"
)

// MPG模式， M： 操作系统的主线程（物理线程）； P：协程执行的上下文环境； G：协程

func main() {
	cpuNum := runtime.NumCPU()

	fmt.Println("cpuNum =", cpuNum)
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("cpuNum =", cpuNum)
}
