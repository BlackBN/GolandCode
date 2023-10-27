package main

import "fmt"

func main() {
	var testMap map[int]string
	fmt.Println(testMap == nil)
	groups := map[string]bool{"": true}
	groups["aaa"] = false
	fmt.Printf("%v\n", groups)

}
