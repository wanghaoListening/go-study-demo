package main

import (
	"flag"
	"fmt"
)

func flagCmd() {

	name := flag.String("name", "david", "default david")
	//使用flag
	flag.Parse()
	fmt.Println(*name)
}
