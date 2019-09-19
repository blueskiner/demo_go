package tcp_comm

import (
	"fmt"
	"net"
)

var (
	serverType    = "tcp"
	serverAddress = "127.0.0.1:9999"
)

func init() {
	fmt.Println("tcp server init")
}

// 启动TCP服务器（内含Loop）
func Start() {
	tcpAddr, err := net.ResolveTCPAddr(serverType, serverAddress)
	if nil != err {
		fmt.Println("net.ResolveTCPAddr error")
		return
	}
	listen, err := net.ListenTCP(serverType, tcpAddr)
	if nil != err {
		fmt.Println("net.ListenTCP error")
		return
	}
	defer listen.Close()

	fmt.Println("start server successful......")

	// 开始接收请求 Loop
	for {
		conn, err := listen.Accept()
		//fmt.Println("accept tcp client %s", conn.RemoteAddr().String())
		if nil != err {
			break
		}
		// 每次建立一个连接就放到单独的协程内做处理
		go handleConnection(conn)
	}
}

func Stop() {
	fmt.Println("退出TCP服务器")
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
}
