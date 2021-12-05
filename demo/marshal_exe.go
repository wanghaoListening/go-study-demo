package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	unMarshalExe()
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func unMarshalExe() {

	str := `{"name":"hhhhhh","age":27}`

	var stu Student

	json.Unmarshal([]byte(str), &stu)

	fmt.Println(stu)
}
