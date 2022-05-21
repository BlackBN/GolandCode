package chapter_eight

import "testing"

func TestStore(t *testing.T){
	ms := Monster{
		name: "adfadfa",
		age: 11,
		kill: 1,
	}
	err := ms.Store()
	if err != nil {
		 t.Fatalf("store is error,%v",err)
	}
}

func TestRestore(t *testing.T){

	msss := Monster{}
	err := msss.ReStore()
	if err != nil {
		t.Fatalf("restore is error,%v",err)
	}
	t.Log(msss.name)
	t.Logf("name %s, age : %d, kill : %d ",msss.name,msss.age,msss.kill)
}
