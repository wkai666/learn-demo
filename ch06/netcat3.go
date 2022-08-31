package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{}) // 定义通道类型
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // 向通道中发送数据
	}()

	mustCopy(conn, os.Stdin) // 发送到服务端
	conn.Close()
	<-done // 阻塞等待从屏幕输出协程中发送的信号
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
