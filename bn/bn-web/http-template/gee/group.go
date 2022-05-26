package gee

import (
	"log"
	"net/http"
	"path"
)

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

// create static handler
func (g *routerGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(g.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		// file := "css/first.css"
		// // Check if file exists and/or if we have permission to access it
		// if _, err := fs.Open(file); err != nil {
		// 	c.Status(http.StatusNotFound)
		// 	return
		// }

		fileServer.ServeHTTP(c.W, c.Req)
	}
}

// serve static files
func (g *routerGroup) Static(relativePath string, root string) {
	handler := g.createStaticHandler(relativePath, http.Dir(root))
	//Register GET handlers
	//加上 ip:port/assets/*filepath 前缀树路由匹配
	g.GET(path.Join(relativePath, "/*filepath"), handler)
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
