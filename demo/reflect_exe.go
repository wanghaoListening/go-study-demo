package main

import (
	"fmt"
	"reflect"
)

type Door struct {
}

func reflectType(x interface{}) {

	v := reflect.TypeOf(x)
	fmt.Printf("the type is : %v\n", v)
	fmt.Printf("the name is : %v , the kind is : %v\n", v.Name(), v.Kind())
}

func main() {

	var name string = "john"
	reflectType(name)

	var age int32 = 10
	reflectType(age)

	var door = Door{}
	reflectType(door)
}
