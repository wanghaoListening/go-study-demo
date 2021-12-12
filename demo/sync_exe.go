package main

import (
	"fmt"
	"sync"
	"time"
)

var x int = 0

var wait sync.WaitGroup

var lock sync.Mutex

var rwLock sync.RWMutex

var loadIconOnce sync.Once

//线程安全的map

var mp sync.Map

func mapExe() {

	for i := 0; i < 10; i++ {
		go func(n int) {
			timeKey := time.Now()
			mp.Store(timeKey, n)
			fmt.Println(mp.Load(timeKey))
		}(i)
	}

}

func loadIcons() {

	//加载图片逻辑
}

//Icon 是并发安全的
func Icon() {
	loadIconOnce.Do(loadIcons)
}

func readValue() {

	rwLock.RLock()
	fmt.Println(x)
	rwLock.RUnlock()
}

func writeValue() {

	rwLock.Lock()
	x += 1
	rwLock.Unlock()
}

func Add() {

	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += i
		lock.Unlock()
	}
	wait.Done()
}

func mainExecAdd() {

	wait.Add(2)
	go Add()
	go Add()
	wait.Wait()
	fmt.Println(x)
}

func main() {
	mainExecAdd()
}
