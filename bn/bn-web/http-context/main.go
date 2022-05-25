package main

import (
	"GolandCode/bn/bn-web/http-context/gee"
	"fmt"
)

func main() {

	e := gee.New()
	e.Get("/", func(c *gee.Context) {
		fmt.Fprintf(c.W, "URL.path = %q\n", c.Path)
	})
	e.Post("/hello", func(c *gee.Context) {
		for headerKey, headerValue := range c.Req.Header {
			fmt.Fprintf(c.W, "Header[%q]=%q\n", headerKey, headerValue)
		}
		fmt.Fprintf(c.W, "URL.path = %q\n", c.Path)
	})
	e.Run("127.0.0.1:1236")
}
