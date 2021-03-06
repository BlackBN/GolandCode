package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	W                 http.ResponseWriter
	Req               *http.Request
	Path              string
	Method            string
	StatusCode        int
	Params            map[string]string
	MiddlewareHandler []HandlerFunc //中间件
	Index             int
}

type H map[string]interface{}

func (c *Context) Param(key string) string {
	value, ok := c.Params[key]
	if !ok {
		return ""
	}
	return value
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W:      w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		Index:  -1,
	}
}

func (c *Context) Next() {
	c.Index++
	for ; c.Index < len(c.MiddlewareHandler); c.Index++ {
		c.MiddlewareHandler[c.Index](c)
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.W.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.W.Write([]byte(html))
}

func (c *Context) Fail(code int, errInfo string) {
	c.Status(code)
	c.W.Write([]byte(errInfo))
}
