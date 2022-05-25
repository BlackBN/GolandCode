package gee

import (
	"fmt"
	"log"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, path string, handler HandlerFunc) {
	if r == nil || r.handlers == nil {
		panic("invalid router")
	}
	log.Printf("Route %4s - %s", method, path)
	r.handlers[method+"-"+path] = handler
}

func (r *router) handle(c *Context) {
	if handler, ok := r.handlers[c.Method+"-"+c.Path]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.W, "404 NOT FOUND : %s|%s\n", c.Method, c.Path)
	}
}
