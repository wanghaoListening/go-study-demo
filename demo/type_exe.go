package main

import "fmt"

type myInt int //自定义类型
type yourInt = int //类型别名

func printType()  {

	var n myInt
	n = 100
	fmt.Println(n)
}

// go 语言函数传参数永远是拷贝



