package main

import (
	"fmt"
	"time"
)

func timeExe() {
	now := time.Now()
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())

	//时间戳
	timestamp := now.Unix()
	fmt.Println(timestamp)
	fmt.Println(now.UnixNano())
	fmt.Println(now)
	fmt.Println(now.Add(24 * time.Hour))

	//2006 12345
	fmt.Println(now.Format("2006-01-02 15:04:05"))

	//定时器
	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t)
	}

	//sleep
	time.Sleep(5)
}
