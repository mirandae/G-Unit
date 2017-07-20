// Author mirandae

package main

import (
	"flag"
	"fmt"
	"os"
)

func getArgs() bool {
	success := flag.Bool("success", false, "its a boolean, dummy")
	flag.Parse()
	return *success
}

func main() {
	if success := getArgs(); success {
		fmt.Println("yay!")
		os.Exit(1)
	}

	fmt.Println("nay")

}
