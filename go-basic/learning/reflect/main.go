package main

import (
	"fmt"
	"reflect"
)

type CreateSchool interface {
	CreateSchool() *School
}

type personalCreateSchool struct {
}

func (p *personalCreateSchool) CreateSchool() *School {
	return &School{
		Name:      "changchunligongdaxue",
		SchoolAge: 100,
	}
}

type User struct {
	School *School
	Name   string
	Age    int16
	Sex    string
}

type School struct {
	Name      string
	SchoolAge int
}

func main() {
	var personalCreateSchool CreateSchool = &personalCreateSchool{}
	user := User{
		Name:   "test",
		Age:    12,
		Sex:    "man",
		School: personalCreateSchool.CreateSchool(),
	}
	typeUser := reflect.TypeOf(user)
	typePersonalCreateSchool := reflect.TypeOf(personalCreateSchool)
	fmt.Println(typeUser.Name(), typeUser.Kind())
	typePersonalCreateSchoolELm := typePersonalCreateSchool.Elem()
	fmt.Println(typePersonalCreateSchool.Name(), typePersonalCreateSchool.Kind())
	fmt.Println(typePersonalCreateSchoolELm.Name(), typePersonalCreateSchoolELm.Kind())
	test(func() {
		fmt.Println("test reflect")
	})
}

func test(inter interface{}) {
	value := reflect.ValueOf(inter)
	interType := reflect.TypeOf(inter)
	//fmt.Println(value.Type(), value.Kind())
	fmt.Println(interType, interType.Kind())
	params := make([]reflect.Value, interType.NumIn())
	// for i := 0; i < interType.NumIn(); i++ {
	// 	params[i] = reflect.ValueOf()
	// }

	value.Call(params)
}
