package main

import (
	"fmt"
	"sort"
)

// func main() {
// 	chan1 := make(chan int)
// 	chan2 := make(chan int)

// 	writeFlag := true
// 	go func() {
// 		for {
// 			if writeFlag {
// 				chan1 <- 1
// 			}
// 			time.Sleep(5 * time.Second)
// 		}
// 	}()
// 	go func() {
// 		for {
// 			if writeFlag {
// 				chan2 <- 1
// 			}
// 			time.Sleep(5 * time.Second)
// 		}
// 	}()

// 	select {
// 	case <-chan1:
// 		fmt.Println("chan1")
// 	case <-chan2:
// 		fmt.Println("chan2")
// 	}
// 	fmt.Println("main exit.")
// }

// func main() {
// 	chan1 := make(chan int)
// 	chan2 := make(chan int)

// 	go func() {
// 		close(chan1)
// 	}()
// 	go func() {
// 		close(chan2)
// 	}()

// 	select {
// 	case <-chan1:
// 		fmt.Println("chan1")
// 	case <-chan2:
// 		fmt.Println("chan2")
// 	}
// 	fmt.Println("main exit.")
// }

// func main() {
// 	//fmt.Println(time.Second << uint(4))
// 	for i := 0; i < 4; i++ {
// 		queryAll()
// 		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
// 	}
// }

// func queryAll() int {
// 	ch := make(chan int)
// 	for i := 0; i < 3; i++ {
// 		go func() { ch <- query() }()
// 	}
// 	return <-ch
// }

// func query() int {
// 	n := rand.Intn(100)
// 	time.Sleep(time.Duration(n) * time.Millisecond)
// 	return n
// }

func main() {
	a := [3]int{1, 3, 4}
	b := [3]int{1, 3, 4}
	fmt.Println(a == b)
	// sd := "tte`2sd';ad;adawe; "
	// fmt.Printf("%T %v\n", sd[18], sd[18])
	// fmt.Println(len(sd))
	// var rmdir []func()
	// s := 0
	// for _, d := range []int{1, 2, 3, 4, 5, 6, 7} {
	// 	// fmt.Printf("d : %d\n", d)
	// 	// fmt.Printf("s : %d\n", s)
	// 	//x := d
	// 	//
	// 	rmdir = append(rmdir, outer(d))
	// 	s++

	// }
	// for _, rm := range rmdir {
	// 	rm()
	// }

	// rmdir[0]()
	// rmdir[0]()
	// rmdir[0]()
	// rmdir[0]()
	// rmdir[0]()
	fmt.Println(threeSum([]int{1, 2, 3, 4, 4, 45, 5, 5, 5, 5, 55, 9}))
	fmt.Println(lengthOfLongestSubstring("dvdf"))

}

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

// 注意：答案中不可以包含重复的三元组。

//

// 示例 1：

// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
// 示例 2：

// 输入：nums = []
// 输出：[]
// 示例 3：

// 输入：nums = [0]
// 输出：[]

//暴力输出
func threeSum(nums []int) [][]int {
	threeMap := make(map[[3]int]bool)
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					sortT := sortThree(nums[i], nums[j], nums[k])
					if _, ok := threeMap[sortT]; !ok {
						threeMap[sortT] = true
					}

				}
			}
		}
	}
	for k, _ := range threeMap {
		data := []int{k[0], k[1], k[2]}
		result = append(result, data)
	}
	return result
}

//
func threeSum2(nums []int) [][]int {
	threeMap := make(map[[3]int]bool)
	result := make([][]int, 0)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)

	for i := 0; nums[i] < 1 && i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		target := 0 - nums[i]
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				sortT := sortThree(nums[i], nums[left], nums[right])
				if _, ok := threeMap[sortT]; !ok {
					threeMap[sortT] = true
				}
				left++
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	for k := range threeMap {
		result = append(result, []int{k[0], k[1], k[2]})
	}
	return result
}

func sortThree(a, b, c int) [3]int {
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	return [3]int{a, b, c}
}

func lengthOfLongestSubstring(s string) int {
	var maxLen int
	hashSet := make(map[uint8]bool)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			_, ok := hashSet[s[j]]
			if ok {
				for k := i; k < j; k++ {
					delete(hashSet, s[k])
				}
				break
			} else {
				hashSet[s[j]] = true
			}
			hashSetLen := len(hashSet)
			if hashSetLen > maxLen {
				maxLen = hashSetLen
			}
		}
	}
	return maxLen
}

// func squares() func() int {
// 	var x int
// 	return func() int {
// 		x++
// 		var y int = 10

// 		return x*x + y
// 	}
// }

// func outer(x int) func() {
// 	f := squares
// 	f()
// 	var y int
// 	y++
// 	return func() {
// 		y++
// 		fmt.Printf("x :%d, y : %d\n", x, y)
// 	}
// }

// func main() {
// 	// f := squares()
// 	// fmt.Println(f()) // "1"
// 	// fmt.Println(f()) // "4"
// 	// fmt.Println(f()) // "9"
// 	// fmt.Println(f()) // "16"
// 	// s := outer()
// 	// s()
// }
