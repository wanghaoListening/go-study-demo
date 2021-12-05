package main

import "fmt"

//结构体占用一块连续的内存空间
type person struct {

	name string
	age int
	hobby []string
}

func structFun()  {

	//结构体指针1
	var p = new(person)
	p.name = "hao"

	//结构体指针2
	var p2 = &person{
		name: "hao",
		age: 27,
		hobby: []string{"篮球","足球"},
	}
	fmt.Println(p2)

	p3 := newPerson("me",19)
	fmt.Println(p3)

}

//构造函数，返回的是结构体还是结构体指针，当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string,age int) person {

	return person{
		name: name,
		age: age,
	}
}


//结构体遇到的问题

