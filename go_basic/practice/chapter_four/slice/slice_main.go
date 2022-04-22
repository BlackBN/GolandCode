package main

import (
	"fmt"
)

func main() {
	fmt.Println("test slice")
	//第一种方式 创建切片，指定长度和容量都为5
	//slice_one_int := make([]int, 5)

	//创建切片，指定长度为3、容量为5的切片
	//slice_one_int := make([]int, 3, 5)

	//第二种方式 创建切片，指定长度和容量相等，都为值的个数
	//slice_two_int := []string{"a", "c", "d", "f"}

	//第三种方式 创建切片，长度和容量为100，第100个初始化为c
	//slice_three_int := []string{99: "c"}

	//第四种方式 创建一个nil切片，长度为0，容量为0
	//var slice_four_int []int

	//fmt.Println(slice_four_int)
	//第四种方式 创建一个空的切片，长度为0，容量为0
	//slice_five_int := make([]int, 0)
	//slice_five_int := []int{}
	//fmt.Println(slice_five_int)

}
