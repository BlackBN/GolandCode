package main

import "fmt"

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
	var rmdir []func()
	s := 0
	for _, d := range []int{1, 2, 3, 4, 5, 6, 7} {
		// fmt.Printf("d : %d\n", d)
		// fmt.Printf("s : %d\n", s)
		//x := d
		//
		rmdir = append(rmdir, outer(d))
		s++

	}
	for _, rm := range rmdir {
		rm()
	}

	rmdir[0]()
	rmdir[0]()
	rmdir[0]()
	rmdir[0]()
	rmdir[0]()

}

func squares() func() int {
	var x int
	return func() int {
		x++
		var y int = 10

		return x*x + y
	}
}

func outer(x int) func() {
	f := squares
	f()
	var y int
	y++
	return func() {
		y++
		fmt.Printf("x :%d, y : %d\n", x, y)
	}
}

// func main() {
// 	// f := squares()
// 	// fmt.Println(f()) // "1"
// 	// fmt.Println(f()) // "4"
// 	// fmt.Println(f()) // "9"
// 	// fmt.Println(f()) // "16"
// 	// s := outer()
// 	// s()
// }
