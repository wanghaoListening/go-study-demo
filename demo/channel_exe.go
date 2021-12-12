package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func channelIni() {

	var channel chan int //需要指定通道类型

	//通道必须初始化
	channel = make(chan int)

	//通道的操作
	//1. 发送 channel <- 1
	//2. 接收 <-channel
	//3. 关闭close()
	wg.Add(1)
	go func() {
		defer wg.Done()
		result := <-channel
		fmt.Println("后台goroutine 从通道当中获取 ", result)
	}()

	channel <- 100
	fmt.Println(channel)
	//关闭通道
	close(channel)
}

//channel练习
//1.启动一个goroutine,生成100个数发送到ch1
//2.启动一个goroutine,从ch1当中取值，计算其平方放到ch2
//3.从ch2取值打印出来

func f1(ch1 chan int) {

	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1) //不能往里写但是可以继续读取
}

func f2(ch1, ch2 chan int) {

	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func changeChannel() {

	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(2)
	f1(a)
	f2(a, b)
	wg.Wait()

	for i := range b {
		fmt.Println(i)
	}
}

//单向通道<-chan ,chan<-

func main() {

	//channelIni()
	changeChannel()
}
