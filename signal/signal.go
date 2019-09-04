package signal

import "fmt"

func init() {
	fmt.Printf("我就是个信号模块的初始化函数\n")
}

func PrintLove() {
	fmt.Printf("SignalPrintLove()\n")
}

func testSignalPrintLove() {
	fmt.Printf("testSignalPrintLove()\n")
}
