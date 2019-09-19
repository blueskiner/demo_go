package main

import (
	"./tcp_comm"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func init() {
	fmt.Printf("main init()\n")
}

func main() {
	fmt.Printf("The operating system is: %s\n", runtime.GOOS)
	//log.SetFlags(log.Ldate | log.Llongfile)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	go watchAllSignals()

	//web.Start()
	tcp_comm.Start()
}

func watchAllSignals() {
	c := make(chan os.Signal) // 监听所有信号
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	s := <-c

	if syscall.SIGHUP == s {
		fmt.Println("收到终端控制进程结束(终端连接断开)消息")
	} else if syscall.SIGINT == s {
		fmt.Println("收到用户发送INTR字符(Ctrl+C)消息")
	} else if syscall.SIGQUIT == s {
		fmt.Println("收到用户发送QUIT字符(Ctrl+/)消息")
	} else if syscall.SIGFPE == s {
		fmt.Println("收到算术运行错误(浮点运算错误、除数为零等)消息")
	} else if syscall.SIGTERM == s {
		fmt.Println("收到结束程序(可以被捕获、阻塞或忽略)消息")
	} else {
		log.Println("other", s)
	}

	tcp_comm.Stop()
}
