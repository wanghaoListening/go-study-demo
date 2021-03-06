package main

import "strings"

// Split abc,b,c ->[abc,b,c]
func Split(str, sep string) []string {

	//预申请内存
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)

	for index > 0 {
		ret = append(ret, str[:index])
		str = str[index+1:]
		index = strings.Index(str, sep)
	}
	ret = append(ret, str)
	return ret
}
