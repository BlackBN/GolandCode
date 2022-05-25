package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	strs := parserPath("/test/sa/s")
	fmt.Println(strs)

	testPrintTrie()
	if !(len(strs) == 3) {
		t.Error("strs leng not equal 4")
	}
}

func testPrintTrie() {
	n := &node{
		part:     "",
		children: make([]*node, 0),
	}

	n.insert("/:/dsafa", parserPath(":/dsafa"), 0)
	n.insert("/sdfa/*fasdfad/", parserPath("/sdfa/*fasdfad/"), 0)
	n.insert("/sdfa/*c/", parserPath("/sdfa/*c/"), 0)
	n.insert("/sdfa/sadfa//", parserPath("/sdfa/sadfa//"), 0)
	n.insert("/:adfadfad/a/b/c", parserPath("/:adfadfad/a/b/c"), 0)
	n.insert("/sdfa/:/:/:/:", parserPath("/sdfa/:/:/:/:"), 0)
	n.insert("/cd/daf/ad/adfa//", parserPath("/cd/daf/ad/adfa//"), 0)
	n.printAllChildString()

}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parserPath("/hello/:name"), []string{"hello", ":name"})
	ok = ok && reflect.DeepEqual(parserPath("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parserPath("/p/*name/*"), []string{"p", "*name", "*"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}
