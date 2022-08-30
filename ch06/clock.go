package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// Listen 函数创建了一个 net.Listener 对象，这个对象会监听一个网络端口上到来的连接
	// 监听 TCP 的 localhost:8000 端口
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// listener对象的Accept方法会直接阻塞，直到一个新的连接被创建，
		// 然后会返回一个net.Conn对象来表示这个连接。
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// 处理一个完整的客户端链接
		go handleConn(conn) // go 使得每一次 handleConn 的调用都进入一个独立的的 goroutine
	}
}

func handleConn(c net.Conn) {
	// 关闭服务侧的连接
	defer c.Close()
	for {
		// 获取当前时刻，写到客户端
		// time.Now().Format() 用来格式化日期和时间信息
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
