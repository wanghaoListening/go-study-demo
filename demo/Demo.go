package main

import (
	"fmt"
	"go/types"
)
import "unsafe"

func mainTo() {
	var name = "cat"
	var age int = 21
	var flag bool = true
	var target = fmt.Sprintf("Hello, %s!, you are %d years old",name,age)
	fmt.Println(target)
	fmt.Println(flag)
	_,number,strs :=numbers()
	fmt.Println(number,strs)
	fmt.Println(area())
	forDemo()
	result := max(10,20)
	fmt.Println(result)
	fmt.Println(swap("google","Runoob"))

	arrayDemo()
	pointerDemo()
	structDemo()
	sliceDemo()
	rangeDemo()
	mapDemo()
	typeConverter()
	implDemo()
	calSum()
	channelBuffer()
}

func numbers()(int,int,string){
	a , b , c := 1 , 2 , "str"
	return a,b,c
}

func area() int{
	const LENGTH int =  10
	const WIDTH int = 5
	area := LENGTH * WIDTH
	return area
}

func ifDemo() int{
	var a int = 20
	if a < 20{
		fmt.Println("a < 20")
	}else {
		fmt.Println("a >=20")
	}
	return a
}

func switchDemo() {

	var grade string = "B"
	var marks int = 90

	switch marks {
		case 90: grade = "A"
		case 80: grade = "B"
		case 50,60,70: grade = "C"
		default: grade = "D"
	}

	switch  {

		case grade == "A":
			fmt.Println("优秀")
	    case grade == "B":
			fmt.Println("良好")
		default:
			fmt.Println("不及格")

	}


	var x interface{}

	switch i := x.(type){

		case types.Nil:
			fmt.Println("x的类型：%T",i)
		case int:
			fmt.Println("x是int 类型")
		default:
			fmt.Println("未知类型")

	}
}

func forDemo() {

	sum :=0
	for i :=0;i<10;i++{
		sum+=i
	}
	fmt.Println(sum)

	initSum :=1
	for initSum <=10 {
		initSum+=initSum
	}
	fmt.Println(initSum)

	//for each range  range 为关键字
	strings := []string{"google","runoob"}
	for i,s :=range strings {
		fmt.Println(i,s)
	}

	numbers := [6]int{1,2,3,4,5}
	for i,x := range numbers{
		fmt.Println(i,x)
	}


}

func max(num1,num2 int) int {

	//声明局部变量
	var result int
	if num1 > num2{
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(x,y string) (string,string) {

	return y,x
}

func swapReference(x *int,y *int)  {

	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func arrayDemo()  {

	//var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)

	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance2)

	//将索引为1和3的元素初始化
	balance3 := [5]float32{1:2,3:7.8}

	fmt.Println(balance3[3])

	//输出元素
	for i:=0;i<len(balance);i++ {
		fmt.Println(balance[i])
	}
}

func pointerDemo()  {

	var a int= 20   /* 声明实际变量 */
	var ip *int        /* 声明指针变量 */

	ip = &a  /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a  )

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	/* 使用指针访问值 */
	fmt.Println(*ip)

	//空指针
	var ptr *int
	if ptr != nil{

		fmt.Println("不是空指针")
	}else {
		fmt.Println("是空指针")
	}
}

func structDemo()  {

	book := Books{"go语言","www.runooob.com","go语言教程",6495407}
	fmt.Println(book)

	var book1 Books

	book1.title = "java教程"
	book1.author = "www.java.com"
	book1.subject = "this java doc"
	book1.bookId = 76543

	fmt.Println(book1)
}

func sliceDemo()  {

	var numbers = make([]int,3,5)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(numbers),cap(numbers),numbers)

	sliceNumbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(sliceNumbers)

	fmt.Println("numbers ==",sliceNumbers)
	fmt.Println("numbers[1:4] ==",sliceNumbers[1:4])
	fmt.Println("numbers[:3] ==",sliceNumbers[:3])
	fmt.Println("numbers[4:] ==",sliceNumbers[4:])

	sliceNumbers1 := make([]int,0,5)
	printSlice(sliceNumbers1)

	sliceNumbers2 := sliceNumbers[:2]
	printSlice(sliceNumbers2)

	sliceNumbers3 := sliceNumbers[2:5]
	printSlice(sliceNumbers3)



}


func printSlice(x []int)  {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

func rangeDemo()  {

	number := []int{2,3,4}
	sum :=0
	for _, num :=range number{
		sum+=num
	}
	fmt.Println(sum)

	//range也可以用在map的键值对上。
	kvs :=map[string]string{"a":"apple","b":"banana"}

	for k,v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func mapDemo()  {

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	for country :=  range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}

	//删除元素
	delete(countryCapitalMap,"France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除后的条目")

	for country :=  range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}

}



func Factorial(n uint64)(result uint64) {
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func typeConverter()  {

	var sum int = 17
	var count int = 5
	var mean = float32(sum)/float32(count)
	fmt.Println(mean)

}


type Phone interface {

	call()
	sendMessage()
}

type Iphone struct {

	name string
	price float64
}

type HuaWei struct {

	name string
	price float64
}

type Books struct {

	title string
	author string
	subject string
	bookId int
}

func (h HuaWei) call() {
	fmt.Println("hua wei ......")
}

func (h HuaWei) sendMessage() {
	fmt.Println("hua wei message")
}

func implDemo()  {

	mate30 :=HuaWei{
		name: "mate 30",
		price: 6999.0,
	}
	mate30.call()
}

func sumChannel(s []int , c chan int)  {

	sum :=0
	for _,v := range s{
		sum+=v
	}
	c <- sum //把sum发送到通道
}

func calSum() {

	s := []int{7,2,8,-9,4,0}
	c := make(chan int)
	go sumChannel(s[:len(s)/2],c)
	go sumChannel(s[len(s)/2:],c)
	x, y := <-c, <-c
	fmt.Println(x,y,x+y)

}

// 通道缓冲区
func channelBuffer()  {

	ch := make(chan int,2)

	//ch是带缓冲区的通道，我们可以同时发送两个数据
	//不需要立刻去同步读取数据
	ch <- 1
	ch <- 2

	//获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//定义枚举,常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数
const (
	UNKNOWN = 0
	FEMALE = 1
	MALE = unsafe.Sizeof("abc")
)

const (
	a = iota
	b
	c
)


