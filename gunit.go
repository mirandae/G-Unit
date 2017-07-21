package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

const (
	passFilename = "output/pass.txt"

	failFilename = "output/fail.txt"
)

var passLyrics []string
var failLyrics []string

// Run loads the ..
func Run() {
	var err error
	passLyrics, err = loadLyrics(passFilename)
	if err != nil {
		log.Fatal("Could not load passing lyrics:", err)
	}
	failLyrics, err = loadLyrics(failFilename)
	if err != nil {
		log.Fatal("Could not load failing lyrics:", err)
	}

	FailTest()
	PassTest()
}

func FailTest() {
	fmt.Println(getRandomLyric(failLyrics))
}

func PassTest() {
	fmt.Println(getRandomLyric(passLyrics))
}

func getRandomLyric(lyrics []string) string {
	return lyrics[rand.Intn(len(lyrics))]
}

func loadLyrics(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failure reading in %s", filename)
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
