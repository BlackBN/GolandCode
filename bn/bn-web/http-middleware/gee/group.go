package gee

import "log"

type routerGroup struct {
	prefix      string
	parent      *routerGroup
	engine      *engine
	middlewares []HandlerFunc
}

func (g *routerGroup) Group(prefix string) *routerGroup {
	e := g.engine
	group := &routerGroup{
		prefix: g.prefix + prefix,
		parent: g,
		engine: e,
	}
	e.groups = append(e.groups, group)
	return group
}

func (g *routerGroup) Use(handler ...HandlerFunc) {
	g.middlewares = append(g.middlewares, handler...)
}

func (g *routerGroup) addRouter(method string, path string, handler HandlerFunc) {
	pattern := g.prefix + path
	log.Printf("Route %4s - %s", method, pattern)
	g.engine.r.addRouter(method, pattern, handler)
}

func (g *routerGroup) GET(prefix string, handler HandlerFunc) {
	g.addRouter("GET", prefix, handler)
}

func (g *routerGroup) POST(prefix string, handler HandlerFunc) {
	g.addRouter("POST", prefix, handler)
}
