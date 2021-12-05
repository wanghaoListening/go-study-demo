package main

import (
	"fmt"
	"os"
)

/**
	学生管理系统
    写一个系统可以查看，新增，删除学生。
 */

type address struct {

	country string
	city string
}

type student struct {

	id int64
	name string
	addr address
}

var (
	allStudent  map[int64]*student
)

func mainStudent() {
	allStudent = make(map[int64]*student,10) //初始化，开辟内存空间
	//1.打印菜单
	fmt.Println("欢迎光临学生管理系统")
	fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
	`)

	for {
		//2.等待用户选择要做什么
		fmt.Println("请输入对应的数字")
		var choice int
		fmt.Scanln(&choice)

		//3.执行对应的函数
		switch choice {

		case 1:
			showAllStudent()
		case 2:
			addNewStudent()
		case 3:
			deleteStudent()
		default:
			fmt.Println("输入有误")
			os.Exit(0)
		}
	}
}

func newStudent(id int64,name string)  *student{

	return &student{
		id: id,
		name: name,
	}
}

func deleteStudent() {

	var id int64
	fmt.Println("请输入编号")
	fmt.Scanln(&id)
	delete(allStudent,id)

}

func addNewStudent() {

	var (
		id int64
		name string
	)
	fmt.Println("请输入编号")
	fmt.Scanln(&id)
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)

	newStu := newStudent(id,name)
	allStudent[id]= newStu
}

func showAllStudent() {

	//把所有的学生打印出来
	for k, v := range allStudent{
		fmt.Printf("学号:%d 姓名:%s\n",k,v.name)
	}
}