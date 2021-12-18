package main

import (
	"fmt"
	"net"
)

func processConn(conn net.Conn) {

	//与客户端通信
	var tmp [128]byte
	n, err := conn.Read(tmp[:])
	if err != nil {
		fmt.Println("read failed")
	}
	fmt.Println(string(tmp[:n]))
}

func serverProcess() {

	//1.监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("server 启动失败")
	}
	for {
		//2.建立连接
		conn, errRes := listener.Accept()
		if errRes != nil {
			fmt.Println("conn failed")
		}
		//3.与客户端通信
		processConn(conn)
	}

}

func clientProcess() {

	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("conn fail")
	}
	var str = "hello ,i am ok"
	conn.Write([]byte(str))
}
