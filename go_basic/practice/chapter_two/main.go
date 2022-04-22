package main

import "fmt"

//import (
//	"fmt"
//	"log"
//	"os"
//
//	_ "webapp/go_in_action/chapter_two/matchers"
//)
//
//func init() {
//	fmt.Println("main init method start")
//	log.SetOutput(os.Stdout)
//}
//
//func main() {
//	// search.Run()
//	//striaa:= "canary-adsfadsfadf"
//	//b := striaa[7:len(striaa)-1]
//	//fmt.Printf("%s",b)
//
//}

var origin = "http://127.0.0.1:8080/"
var url = "ws://127.0.0.1:9000/ws"

type A struct {
	name string
	age int
}

type B struct {
	info string
	A map[string]*A
}

func main() {
	b := B{
		info: "test",
		A:    make(map[string]*A),
	}
	a := A{
		name: "xujiayu",
		age:  12,
	}

	b.A["ond"]=&a
	b.A["onf"]=&a
	b.A["ong"]=&a
	fmt.Println(b.A["ond"].name,b.A["onf"].name,b.A["ong"].name)

	delete(b.A,"ond")
	fmt.Println(b.A["onf"].name,b.A["ong"].name)

	//ws, err := websocket.Dial(url,"",origin)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//message := []byte("hello, world!你好")
	//_, err = ws.Write(message)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Send: %s\n", message)
	//
	//var msg = make([]byte, 512)
	//m, err := ws.Read(msg)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Receive: %s\n", msg[:m])
	//
	//ws.Close()//关闭连接
}
