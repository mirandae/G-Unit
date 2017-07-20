package main

import (
	"flag"
	"fmt"
)

func getArgs() bool {
	success := flag.Bool("success", false, "just use it correctly")
	flag.Parse()
	return *success
}

func main() {
	success := getArgs()

	if success {
		fmt.Println("yay!")
		return
	}
	fmt.Println("nay")
}
