package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type A struct {
	mapinfo map[string]*B
}

type B struct {
	name string `json:"name"`
	age int `json:"age"`
	c *C `json:"c"`
	d D `json:"d"`
}

type C struct {
	message string `json:"message"`
}

type D struct {
	info string `json:"info"`
}

func test(){
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(i)
	}
}

func main()  {
	//go test()
	//time.Sleep(6000)
	b := &B{
		name: "a",
		age: 12,
		c: &C{
			message: "aaa",
		},
		d: D{
			info: "aa",
		},
	}
	mapInf := make(map[int]*B)
	for i := 0 ; i < 10 ; i++ {
		mapInf[i] = b
	}
	for key, value := range mapInf {
		fmt.Printf("key info %d" , key)
		fmt.Println()
		if key == 5 {
			value.age = 40
		}
		fmt.Println(value.age)
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		//fmt.Println(f.Name())
		if f.IsDir() {
			fss ,_ := ioutil.ReadDir(getFilePath(".",f.Name()))
			for _, fs := range fss {
				fmt.Println(fs.Name())
			}
		}
	}


	fmt.Println("-----------------")
	fmt.Println(getFilePath("a","b","c"))
	fmt.Println(getFilePath("aaaa","ddd","s","sdfa","asdfasd","Asdfasd.adsfa"))
}

func getFilePath(dirNames ...string) string {
	path := ""
	if dirNames == nil || len(dirNames) < 1 {
		return path
	}
	for _, dirName := range dirNames {
		path = path + dirName + "/"
	}
	return path[0 : len(path)-1]
}
