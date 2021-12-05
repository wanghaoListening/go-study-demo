package main

import "fmt"

func scanFunc()  {
	var s string
	fmt.Scan(&s)
	fmt.Println(s)

	var(
		name string
		age int
		class string
	)

	fmt.Scanf("%s %d %s",&name,&age,&class)


}
