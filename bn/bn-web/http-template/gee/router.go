package gee

import (
	"fmt"
)

type router struct {
	handlers    map[string]HandlerFunc
	routerNodes map[string]*node
}

func newRouter() *router {
	return &router{
		handlers:    make(map[string]HandlerFunc),
		routerNodes: make(map[string]*node),
	}
}

func (r *router) addRouter(method string, path string, handler HandlerFunc) {
	if r == nil || r.handlers == nil {
		panic("invalid router")
	}
	if _, ok := r.routerNodes[method]; !ok {
		r.routerNodes[method] = &node{}
	}
	r.routerNodes[method].insert(path, parserPath(path), 0)
	r.handlers[method+"-"+path] = handler
}

func (r *router) findRoute(method string, path string) *node {
	node, ok := r.routerNodes[method]
	if !ok {
		return nil
	}
	finalChildNode := node.search(parserPath(path), 0)
	if finalChildNode != nil {
		return finalChildNode
	}
	return nil
}

func (r *router) handle(c *Context) {
	node := r.findRoute(c.Method, c.Path)
	if node != nil {
		handler := r.handlers[c.Method+"-"+node.pattern]
		if handler != nil {
			c.MiddlewareHandler = append(c.MiddlewareHandler, handler)
		}
	} else {
		c.MiddlewareHandler = append(c.MiddlewareHandler, func(c *Context) {
			fmt.Fprintf(c.W, "404 NOT FOUND : %s|%s\n", c.Method, c.Path)
		})
	}
	c.Next()
}
