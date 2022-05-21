package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Animal struct {
	name string
	age int32
}

type User struct {
	name string
	age int32
	color string
	animal Animal
}

type UserCopy User
type Duration int64
type Ducopy int64

func (anima *Animal) GetAnimalInfo(){
	fmt.Println(fmt.Printf("%v",anima))
	anima.name= "test"
	anima.age=5
	fmt.Println(fmt.Printf("in func %v",anima))
}

func main(){
	//var animal Animal
	//animal.name = "little cat"
	//animal := &Animal{
	//	name: "little cat ",
	//	age:  3,
	//}
	//animal.GetAnimalInfo()
	//fmt.Println(fmt.Printf(" out func %v",animal))
	//animal := Animal{"little cat ", 4}
	//fmt.Printf("name : %s, and age : %d",animal.name, animal.age)
	//fmt.Println()
	//user := User{
	//	name:   "test",
	//	age:    0,
	//	color:  "black",
	//	animal: Animal{
	//		name:"little cat",
	//		age:1,
	//	},
	//}
	//var userCopy UserCopy
	//userCopy = UserCopy(user)
	//fmt.Printf("%v",userCopy)
	//fmt.Println()
	//var dur Duration
	//ducopy := Ducopy(12)
	//
	//dur = Duration(ducopy)
	//fmt.Printf("%v   ",dur)

	//ExecuteCurl()
	//
	// TestIoOne() an := Animal{}
	//test:=TestInterBody{
	//
	//}
	//testTwo := TestBody{
	//
	//}
	//test.GetInfo([]byte("hello"))
	//
	//fmt.Println("")
	//Test(testTwo)
	//Test(&test)
	//dur(9).GetInfo([]byte("uu"))
	str :=ToString(Binary(30))
	fmt.Printf("%s",str)
	fmt.Println()
}
type Stringer interface {
	String() string
}

func ToString(str interface{}) string{
	if v,ok := str.(Stringer) ; ok{
		return v.String()
	}
	switch v := str.(type) {
		case int:
			return strconv.Itoa(v)
		case float64:
			return strconv.FormatFloat(v,'E',-1,64)
	}
	return "???"
}

type Binary uint64

func (i Binary)String() string{
	return string(i.Get())
}
func (i Binary) Get() uint64 {
	return 4505
}






type dur int

func (dur *dur)GetInfo(p []byte)(n int,err error){
	return 9,nil;
}


func Test(inter TestInterfaceOne)(){
	vaul,_:=inter.GetInfo([]byte("world"))
	fmt.Printf("ddddd : %v",vaul)
}


func TestIoOne(){
	var bytes bytes.Buffer
	bytes.Write([]byte("hello"))

	fmt.Fprintf(&bytes," world")

	io.Copy(os.Stdout,&bytes)
}

func ExecuteCurl(){
	r,err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Errorf("%v",err)
	}
	io.Copy(os.Stdout,r.Body)
	if err = r.Body.Close() ; err != nil{
		log.Printf("%v",err)
	}

}
type TestBody struct {

}
func (test TestBody) GetInfo(p []byte)(n int,err error){
	fmt.Printf("test get info")
	return 9,nil
}

type TestInterBody struct {


}

func  (body *TestInterBody) GetInfo(p []byte)(n int,err error){
	fmt.Printf("test get info")
	return 9,nil
}

type TestInterface interface {
	GetInfo(p []byte) (n int, err error)
}

type TestInterfaceOne interface {
	GetInfo(p []byte) (n int, err error)
}

type TestInterTwo interface {
	GetInfo(p []byte) (n int, err error)
}

type TestInterThree interface {
	GetInfo(p []byte) (n int, err error)
}



