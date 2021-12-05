package main

import (
	"fmt"
	"strings"
)

//go 语言字符串是用utf-8进行编码的默认支持中文字符串

func stringFunc() {

	//多行字符串
	poem := `
		1.在天愿作比翼鸟,在地愿为连理枝。
		2.心似双丝网,中有千千结。
		3.身无彩凤双飞翼,心有灵犀一点通。
		4.东边日出西边雨,道是无晴却有晴。
	`
	fmt.Println(poem)

	//字符串相关操作
	fmt.Println(len(poem))
	ret := strings.Split(poem," ")
	fmt.Println(ret)
	fmt.Println(strings.Contains(poem,"灵犀"))

	s := "hello你好"

	for _,ch := range s{
		fmt.Printf("%c\n",ch)
	}

	ss := "白萝卜"
	// rune 本质上是一个int32
	sss := []rune(ss)
	fmt.Println(sss)
	fmt.Println(string(sss))
}

