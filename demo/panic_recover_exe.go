package main

import "fmt"

func panicFunc()  {

	defer func() {
		err := recover() //1。recover必须搭配defer使用 2。defer一定要在Panic之前定义
		if err != nil{
			fmt.Println(err)
		}
	}()
	panic("数据库连接失败")//程序崩溃退出
}




