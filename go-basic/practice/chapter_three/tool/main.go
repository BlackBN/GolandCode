package main

import "fmt"

func main() {
	n, err := fmt.Printf("The quick brown fox jumps over the lazy dog,%f", 3.14)
	if n == 0 {
		fmt.Print("n == 0")
	}
	if err != nil {
		fmt.Print("err != nil")
	}
	fmt.Println()
}
