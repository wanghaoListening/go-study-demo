package main

import "fmt"

func deferDemo()  {
	fmt.Println("start")
	fmt.Println("defer ------------") //defer延迟执行，推迟到函数即将返回的时候再去执行
	fmt.Println("end")
}


