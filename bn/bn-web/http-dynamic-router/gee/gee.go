package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type engine struct {
	r *router
}

func New() *engine {
	return &engine{
		r: newRouter(),
	}
}

func (e *engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

// 实现 http.Handler 接口的方法
func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	e.r.handle(newContext(w, req))
}

func (e *engine) Get(path string, handler HandlerFunc) {
	e.r.addRouter("GET", path, handler)
}

func (e *engine) Post(path string, handler HandlerFunc) {
	e.r.addRouter("POST", path, handler)
}

func (e *engine) PrintTrie() {
	router := e.r
	if router != nil {
		routerNodes := router.routerNodes
		for k, v := range routerNodes {
			fmt.Printf("%s method node info\n", k)
			v.printAllChildString()
		}
	}
}
