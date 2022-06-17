package main

import (
	"GolandCode/go-basic/learning/method/point"
	"fmt"
)

func main() {
	p := point.New(10, 20)
	p.Change(11, 21)
	fmt.Printf("%v\n", p)
	p.Change2(12, 22)
	fmt.Printf("%v\n", p)
	var a point.Action = point.Person{}
	fmt.Printf("%T\n", a)
	a.Fight()
}
