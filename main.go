// Author mirandae
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/g-unit/gunit"
)

const (
	gunitStartMenu = "gunit_ascii.txt"
)

func getArgs() bool {
	success := flag.Bool("success", false, "its a boolean, dummy")
	flag.Parse()
	return *success
}

func main() {
	printMenu()

	if success := getArgs(); success {
		fmt.Println("yay!")
		return
	}

	gunit.Run()

}

func printMenu() {
	data, err := ioutil.ReadFile(gunitStartMenu)
	if err != nil {
		log.Fatal("Can't start G-unit :(")
	}
	fmt.Print(string(data))

}
