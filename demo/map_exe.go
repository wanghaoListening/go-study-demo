package main

import "fmt"

func mapFunc()  {

	m1 := make(map[string]string,10)
	m1["alibaba"] = "马云"
	m1["baidu"] = "李彦宏"

	for k, v := range m1 {
		fmt.Println(k,v)
	}

	delete(m1,"baidu")
}
