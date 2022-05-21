// Sample program to show how to show you how to briefly work with io.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// main is the entry point for the application.
func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)

	count := CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}

// CountWords counts the number of words in the specified
// string and returns the count.
func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}

type TestInterfaceTwo interface {
	GetInfo(p []byte) (n int, err error)
}
