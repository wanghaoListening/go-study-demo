package main

import "fmt"

//包名的导入是从go path src 目录下开始的

///interface{} //空接口，所有类型都实现的空接口


//同一个结构体可以实现多个接口
//接口可以嵌套

type animal interface {
	speaker
	mover
}

type speaker interface {
	speak()
}

type mover interface {
	move()
}

type cat struct {

}

type duck struct {

}

func (c cat) speak()  {

	fmt.Println("喵喵")
}

func (c cat) move()  {
	fmt.Println("run run")
}

func (d duck) speak()  {

	fmt.Println("呱呱")
}

func talk(s speaker)  {
	s.speak()
}

func assertType(i interface{})  {

	switch i.(type) {

		case string:
			fmt.Println("是一个字符串")
		case int:
			fmt.Println("是一个树数值")
		default:
			fmt.Println("是其他类型")
	}
}


