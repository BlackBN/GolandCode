package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(c *Context)

type engine struct {
	r      *router
	groups []*routerGroup //如果实现简单的group,这个可以不需要
	*routerGroup
}

func New() *engine {
	engine := &engine{r: newRouter()}
	firstGroup := &routerGroup{engine: engine}

	engine.routerGroup = firstGroup

	engine.groups = append(engine.groups, firstGroup)
	return engine
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
