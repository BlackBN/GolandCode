package gee

import "log"

type routerGroup struct {
	prefix      string
	parent      *routerGroup
	middlewares []HandlerFunc //中间件，后续有用到
	engine      *engine
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
