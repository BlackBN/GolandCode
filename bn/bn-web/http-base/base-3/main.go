package main

import (
	"GolandCode/bn/bn-web/http-base/base-3/gee"
	"fmt"
	"net/http"
)

func main() {
	e := gee.New()
	e.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
	})
	e.Post("/hello", func(w http.ResponseWriter, r *http.Request) {
		for headerKey, headerValue := range r.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", headerKey, headerValue)
		}
		fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
	})
	e.Run("127.0.0.1:1236")
}
