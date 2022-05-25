package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type engine struct {
	router map[string]HandlerFunc
}

func New() *engine {
	return &engine{
		router: make(map[string]HandlerFunc),
	}
}

func (e *engine) addRouter(method string, path string, handler HandlerFunc) {
	if e == nil || e.router == nil {
		panic("invalid engine")
	}
	e.router[method+"-"+path] = handler
}

func (e *engine) Get(path string, handler HandlerFunc) {
	e.addRouter("GET", path, handler)
}

func (e *engine) Post(path string, handler HandlerFunc) {
	e.addRouter("POST", path, handler)
}

func (e *engine) Run(addr string) {
	http.ListenAndServe(addr, e)
}

// 实现 http.Handler 接口的方法
func (e *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	url := req.Method + "-" + path
	if handlerFunc, ok := e.router[url]; ok {
		handlerFunc(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", path)
	}
}
