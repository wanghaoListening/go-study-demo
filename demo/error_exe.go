package main

import (
	"fmt"
	"strings"
)

func getPodName() (string, error) {

	return "", fmt.Errorf("not found pod")
}

func main() {

	addr, err := getPodName()

	if err != nil && err.Error() == "not found pod" {
		fmt.Println("hhhhh")
	} else {
		fmt.Println(addr)
	}
	str := "10.13.111.0"
	fmt.Printf(transformIpv4(str))

}

func transformIpv4(host string) string {
	const count = 3
	arrList := strings.Split(host, ".")
	start := "n" + arrList[1]
	var middle string
	for i := 1; i <= count-len(arrList[2]); i++ {
		middle = middle + "0"
	}
	middle += arrList[2]
	var end string
	for i := 1; i <= 2; i++ {
		end = end + "0"
	}
	end += arrList[3]
	return fmt.Sprintf("%s-%s-%s", start, middle, end)
}
