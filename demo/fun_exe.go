package main

import (
	"fmt"
	"math"
)

func add(a,b int) int {

	return a+b
}


var addFun = func(a,b int) int{
	return a + b
}

func swapFun(a,b int) (int,int){

	return b,a
}

func sum(a int, more ...int) int {

	for _,v := range more{

		a+=v
	}
	return a
}

// 不仅函数的参数可以有名字，也可以给函数的返回值命名
func find( m map[int]int,key int)  (value int,ok bool){

	value,ok = m[key]
	return
}

func deferFunc()  {

	//通过defer语句在return语句之后修改返回值,在for循环内部执行defer语句并不是一个好的习惯，此处仅为示例，不建议使用。
	for i :=0;i<3;i++{
		defer func() {println(i)}()
	}
}

func floatFunc() float64 {

	return math.MaxFloat64
}


//函数也可以作为参数传入
func funcParam(fn func() float64)  {

	ret := fn()
	fmt.Println(ret)
}

//函数也可以作为返回值
func funcReturn(fn func() float64) func(int,int) int{

	ret := func(a,b int) int{

		return a+b
	}
	return ret
}

func annoyFunc()  {

	fn := func(a,b int) {

		fmt.Println(a,b)
	}
	fn(12,3)

	// 如果只调用一次的函数
	func(){
		fmt.Println("hello world")
	}()//后面的括号代表执行函数
}


// go build /go run / go install


//批量声明变量,go语言变量声明了必须使用

var (
	name string
	age int
	isOk bool
)

const(
	//iota是go语言的常量计数器
	aac = iota
	abc = 100
	acc = iota
)

const(
	_ = iota
	KB = 1 <<(10*iota)
	MB = 1 <<(10*iota)
	GB = 1 <<(10*iota)
	TB = 1 <<(10*iota)
	PB = 1 <<(10*iota)
)




