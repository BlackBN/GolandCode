package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}
type Test struct {
	InfoMap map[string]string
}
type Cat struct {
	Name string
	Age  int
}

var (
	ResultMap = make(map[int]int)
	lock      sync.Mutex
)

func TestS(num int) {
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	lock.Lock()
	ResultMap[num] = res
	lock.Unlock()
}

func main() {
	// cpuNumber := runtime.NumCPU()
	// println(cpuNumber)
	// runtime.GOMAXPROCS(cpuNumber - 1)
	// println("AA")
	// for i := 0; i < 200; i++ {
	// 	go TestS(i + 1)
	// }
	// time.Sleep(10 * time.Second)
	// lock.Lock()
	// fmt.Println(ResultMap)
	// lock.Unlock()
	// for i := 0; i < 1000; i++ {
	// 	go TestCallUrl(i)
	// }
	//time.Sleep(60 * time.Second)

	myChan := make(chan int, 3)
	//fmt.Printf("%v %v\n", myChan, &myChan)
	myChan <- 1
	myChan <- 2
	myChan <- 3

	fmt.Printf("%v %v\n", len(myChan), cap(myChan))
	a := <-myChan
	fmt.Println(a)
	fmt.Printf("%v %v\n", len(myChan), cap(myChan))

	myChan <- 4

	//b := <-myChan
	//c := <-myChan
	//fmt.Printf("%d,%d\n", b, c)
	//d := <-myChan
	//fmt.Printf("%d\n", d)

	inteChan := make(chan interface{}, 10)
	inteChan <- Cat{Name: "aa", Age: 12}
	close(myChan)
	for c := range myChan {
		fmt.Printf("c:%d\n", c)
	}

	//go TestCallUrl()
	//go TestCallUrl()
	close(inteChan)
	cat := <-inteChan
	fmt.Printf("%T, %v\n", cat, cat)
	fmt.Printf("%v\n", cat.(Cat).Name)
	//inteChan <- 12
	fmt.Printf("%d\n", time.Now().Unix())

}

func TestCallUrl(i int) {

	req, err := http.NewRequest("GET", "http://127.0.0.1:6700/getAgeInfo", nil)
	// 比如说设置个token
	//req.Header.Set("X-RPC-header", fmt.Sprintf("%d", i))

	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
