package main

import "fmt"

//go 语言中如果标示符首字母大写，标示对外部可见（public，公有的）

//方法和接收者

type car struct {
	name string
	loc string
}

//构造函数
func newCar(name,loc string) car {

	return car{
		name: name,
		loc: loc,
	}
}

//方法是作用于特定类似的函数
//(c car) 表示的是具体的类型
func (c car) run()  {

	fmt.Println("s%:正在跑",c.name)
}

func methodFunc()  {

	c := newCar("奔驰","德国")
	c.run()
}

type myFloat float64

func (m myFloat) hello()  {

	fmt.Println("hello")
}

