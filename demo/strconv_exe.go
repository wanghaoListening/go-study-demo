package main

import (
	"fmt"
	"strconv"
)

func strconvFunc() {

	//从字符串转换成数字
	str := "10001"
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v : %T\n", ret, ret)

	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v : %T\n", retInt, retInt)

	//从字符串转换成布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)

	fmt.Printf("%#v : %T\n", boolValue, boolValue)

	//从字符串转换成浮点数
	floatStr := "3.14"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)

	fmt.Printf("%#v : %T\n", floatValue, floatValue)

}

/*func main() {
	strconvFunc()
}*/
