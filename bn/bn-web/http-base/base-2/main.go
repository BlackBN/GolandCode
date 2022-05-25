package main

import (
	"fmt"
	"net/http"
)

type definitionHttpHandler struct {
}

func NewDefinitionHttpHandler() definitionHttpHandler {
	return definitionHttpHandler{}
}

func (d definitionHttpHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
	case "/hello":
		for headerKey, headerValue := range req.Header {
			fmt.Fprintf(w, "Header[%q]=%q\n", headerKey, headerValue)
		}
		fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL.Path)
	}
}

func main() {
	http.ListenAndServe(":1235", NewDefinitionHttpHandler())
}
