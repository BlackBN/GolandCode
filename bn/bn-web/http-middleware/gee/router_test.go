package gee

import (
	"fmt"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRouter("GET", "/", nil)
	r.addRouter("GET", "/hello/:name", nil)
	r.addRouter("GET", "/hello/b/c", nil)
	r.addRouter("GET", "/hi/:name", nil)
	r.addRouter("GET", "/assets/*filepath", nil)
	return r
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n := r.findRoute("GET", "/hello/geektutu")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	// if ps["name"] != "geektutu" {
	// 	t.Fatal("name should be equal to 'geektutu'")
	// }

	fmt.Printf("matched path: %s, params['name']: \n", n.pattern) //, ps["name"])

}
