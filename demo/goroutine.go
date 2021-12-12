package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

func goHello() {

	fmt.Println("hello go")
}

/*func main() {
	//go goHello() //go 关键字开启一个goroutine

	fmt.Println("hello yu")
	//main 函数结束后，由main启动的goroutine也有结束

	waitGroupDemo()

	//runtime 相关
	procExe()
}*/

//wait group

var wc sync.WaitGroup

func waitDemo() {
	defer wc.Done() // wc 计数器-1
	fmt.Println(rand.Int())
}

func waitGroupDemo() {

	for i := 0; i < 5; i++ {
		wc.Add(1) // wc 计数器+1
		go waitDemo()
	}
	wc.Wait() // 等待wc 计数器减到0
}

//Go GMP原理；goroutine调度模型

//Go GOMAXPROCS

func procExe() {

	runtime.GOMAXPROCS(2)
	fmt.Println(runtime.NumCPU())

}

//m:n把m个goroutine分配给n个操作系统去执行
