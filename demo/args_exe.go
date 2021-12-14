package main

import (
	"fmt"
	"os"
)

func readCommand() {

	if len(os.Args) <= 0 {
		fmt.Println("no args")
	} else {
		fmt.Println(os.Args[0])
	}
}
