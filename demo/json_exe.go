package main

import (
	"encoding/json"
	"fmt"
)

type dog struct {

	Name string `json:"name"`
	Age uint `json:"age"`
}


func jsonFunc()  {


	d := dog{
		Name: "tom",
		Age: 10,
	}

	str,err := json.Marshal(d)
	if err != nil{
		fmt.Println("序列化失败")
	}
	fmt.Println(string(str))

	jsonStr := `{"name":"john","age":11}`

	var d2 dog
	json.Unmarshal([]byte(jsonStr),&d2) //传指针为了能够在json.Unmarshal内部进行修改
	fmt.Println(d2)
}


