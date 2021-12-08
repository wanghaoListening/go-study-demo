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

func reflectVaklue(x interface{}) {

	v := reflect.ValueOf(x)
	//值的种类
	switch v.Kind() {

	case reflect.Int64:
		fmt.Println()

	case reflect.Float64:
		fmt.Println()
	case reflect.Uint32:

	}
}

//通过反射设置变量的值

func reflectSetValue1(x interface{}) {

	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Int64 {
		v.SetInt(300) //修改的是副本
	}
}

func reflectSetValue2(x interface{}) {

	v := reflect.ValueOf(x)

	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(300)
	}
}

func main() {

	var name string = "john"
	reflectType(name)

	var age int32 = 10
	reflectType(age)

	var door = Door{}
	reflectType(door)

	var temp int64 = 20

	reflectSetValue1(&temp)
	fmt.Println(temp)

	reflectSetValue2(&temp)
	fmt.Println(temp)
}
