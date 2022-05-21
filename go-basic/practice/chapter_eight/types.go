package chapter_eight

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	name string `json:"name"`
	age int `json:"age"`
	kill int 	`json:"kill"`
}

func (m *Monster)Store() (err error) {
	data , err := json.Marshal(m)
	if err != nil {
		fmt.Printf("marshal is err")
		return
	}
	err = ioutil.WriteFile("/Users/bn/abc",data,777)
	if err != nil {
		return
	}
	return
}

func (m *Monster)ReStore()(err error){

	data , err := ioutil.ReadFile("/Users/bn/abc")
	if err != nil {
		return
	}
	if err = json.Unmarshal(data,m); err != nil {
		return
	}
	return
}
