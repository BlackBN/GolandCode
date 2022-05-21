package main

import (
	"fmt"
)

func main() {
	ArrayTest()
	//ArrayPointTest()
}

func ArrayTest() {
	fmt.Println("array test begin")

	//第一种方式 默然初始化为0
	var array [5]int
	for i := 0; i < 5; i++ {
		fmt.Printf("%d, ", array[i])
	}
	fmt.Println()

	//第二种方式 指定数组大小，并初始化一些默认值
	array_copy_one := [5]int{10, 20, 30}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d, ", array_copy_one[i])
	}
	fmt.Println()

	//第三种方式 不指定数组大小，初始化的默认值的个数就是该数组的长度
	array_copy_two := [...]int{10, 20, 30}
	for i := 0; i < 3; i++ {
		fmt.Printf("%d, ", array_copy_two[i])
	}
	fmt.Println()

	//第四种方式 指定数组大小，指定某个索引的值
	array_copy_three := [5]int{1: 20, 3: 50}
	for i := 0; i < 5; i++ {
		fmt.Printf("%d, ", array_copy_three[i])
	}
	fmt.Println()

	fmt.Println("array point test begin")

	//声明一个指针数组
	array_point := [5]*int{new(int), new(int), new(int), new(int), new(int)}
	*array_point[0] = 20
	*array_point[4] = 55
	for i := 0; i < 5; i++ {
		fmt.Printf("%d, ", *array_point[i])
	}
	fmt.Println()

	//数组复制
	var array1 [5]string

	array2 := [5]string{"a", "b", "b"}

	array1 = array2

	for i := 0; i < 5; i++ {
		fmt.Printf("%s, ", array1[i])
	}
	fmt.Println()

	//指针数组复制
	var array3 [5]*string
	array4 := [5]*string{new(string), new(string), new(string), new(string), new(string)}
	array3 = array4 //同时指向同一块地址
	*array4[4] = "four"
	for i := 0; i < 5; i++ {
		fmt.Printf("%s, ", *array3[i])
	}
	fmt.Println()

	//多维数组
	//var doubleArray [4][2]int

	//二维数组：{{第一行},{第二行},{第三行}}
	//doubleArray_two := [4][3]int{{10, 0, 23}, {20, 21}, {40, 41}}

	//可以指定某个索引的值
	doubleArray_two := [4][3]int{1: {2: 10}, 0: {1, 20, 30}, 3: {1: 41}}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", doubleArray_two[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	//索引复制与一维数组相似 ， 可以独立复制某个维度的值

	//在函数间传递数组 1000000 * 4 B
	var arrayMethod []int //占4 MB内存
	foo(arrayMethod)
	fooPoint(&arrayMethod)

}

//会在方法栈里面在创建一个 4 MB的内存
func foo(array []int) {
	fmt.Println("is success")
}

//仅仅花费一个指针大小的的内存
func fooPoint(array *[]int) {

}
