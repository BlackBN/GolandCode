package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {
	var a []string
	a = nil
	if len(a) == 0 {
		fmt.Println("len a is zero")
	}

	sliceA := make([]int, 4, 10)
	sliceA[2] = 5
	fmt.Println(sliceA)

	// nil切片 和 空切片
	var s1 []int
	// new 函数返回是指针类型，所以需要使用 * 号来解引用
	var s4 = *new([]int)

	var s2 = []int{}
	var s3 = make([]int, 0)

	fmt.Println(len(s1), len(s2), len(s3), len(s4))
	fmt.Println(cap(s1), cap(s2), cap(s3), cap(s4))
	fmt.Println(s1, s2, s3, s4)

	var a1 = *(*[3]int)(unsafe.Pointer(&s1))
	var a2 = *(*[3]int)(unsafe.Pointer(&s2))
	var a3 = *(*[3]int)(unsafe.Pointer(&s3))
	var a4 = *(*[3]int)(unsafe.Pointer(&s4))
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	// end

	// 切片截取
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 := slice[2:5]
	slice2 := slice1[2:6:7]

	slice2 = append(slice2, 100)
	slice2 = append(slice2, 200)

	slice1[2] = 20

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice)
	// end
	// printCap()
	// printCap2()
	oldSlice := []int{9, 4, 5, 6, 7, 2}
	fmt.Printf("%p\n", oldSlice)
	oldSliceInfo := *(*[3]int)(unsafe.Pointer(&oldSlice))
	fmt.Printf("%x\n", oldSliceInfo[0])
	fmt.Printf("%p\n", &oldSlice)
	testSlice(oldSlice)
	user := User{
		name: "aaa",
		age:  12,
	}
	fmt.Printf("%p\n", &user)
	printStruct(&user)

	var testS []string = make([]string, 0)
	var testX []string
	testB, _ := json.Marshal(testS)
	testC, _ := json.Marshal(testX)
	fmt.Printf("%s\n", testB)
	fmt.Printf("%s\n", testC)
}

type User struct {
	name string
	age  int
}

func printStruct(s *User) {
	fmt.Printf("%v\n", s)
	fmt.Printf("%p\n", s)
}

func testSlice(s []int) {
	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", &s)
	for _, data := range s {
		data++
	}
	fmt.Printf("%v\n", s)
}

// func printCap() {
// 	s := make([]int, 0)

// 	oldCap := cap(s)

// 	for i := 0; i < 4096; i++ {
// 		s = append(s, i)

// 		newCap := cap(s)

// 		if newCap != oldCap {
// 			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i-1, oldCap, i, newCap)
// 			oldCap = newCap
// 		}
// 	}
// }

// func printCap2() {
// 	s := []int{1, 2, 4}
// 	x := []int{4, 5, 6, 4}
// 	s = append(s, x...)
// 	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
// 	s = append(s, []int{3, 2, 3, 4, 4, 5, 5, 5, 6, 6}...)
// 	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
// }

// sliceB := sliceA[2:4]
// sliceA[2] = 29
// fmt.Printf("%v\n", sliceA)
// fmt.Printf("%v\n", sliceB)
// fmt.Printf("%p,%p\n", &sliceA[2], &sliceB[0])
//sliceC := sliceA[:]
// fmt.Println(sliceA)
// fmt.Printf("%p,%p\n", &sliceA, &sliceA[0])
// sliceA = append([]int{32}, sliceA...)
// fmt.Println(sliceA)
// fmt.Printf("%p,%p\n", &sliceA, &sliceA[1])

// sliceC := new([]int)
// fmt.Printf("%v\n", sliceC)
// fmt.Println(sliceC == nil)
