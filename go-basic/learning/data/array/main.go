package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	a[0] = 1
	b := [4]int{1, 3}
	c := [4]int{1: 2, 3: 2}
	d := [...]int{1, 3, 3, 3}
	fmt.Printf("%v,%v,%v,%v\n", a, b, c, d)
	b[1] = 2
	b[2] = 3
	b[3] = 4
	fmt.Println(c == b)
	fmt.Printf("%p,%p,%p,%p\n", &a, &a[0], &a[1], &a[2])
	copy(a)
}

func copy(a [4]int) {
	fmt.Printf("%p,%p,%p,%p\n", &a, &a[0], &a[1], &a[2])

}
