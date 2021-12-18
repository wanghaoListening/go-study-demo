package main

import (
	"fmt"
	"net"
	"strings"
)

//udp server

func UdpServer() {

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})

	if err != nil {
		fmt.Println("listen udp failed ,err:", err)
		return
	}
	defer conn.Close()
	var data [1024]byte

	for {
		n, addr, err := conn.ReadFromUDP(data[:])

		if err != nil {
			fmt.Println("read data failed,the error:", err)
		}
		upperStr := strings.ToUpper(string(data[:n]))

		conn.WriteToUDP([]byte(upperStr), addr)
	}
}

func UdpClient() {

	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {

		fmt.Println("dial udp failed ", err)
	}
	defer conn.Close()

	var data [1024]byte
	var reply [1024]byte
	for {
		//写入数据
		conn.Write(data[:])
		//读取数据
		n, _, err := conn.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("read data failed ", err)
		}
		fmt.Println("收到数据", string(reply[:n]))
	}

}
