package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":1234", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for headerKey, headerValue := range req.Header {
		fmt.Fprintf(w, "Header[%q]=%q\n", headerKey, headerValue)
	}
	fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
}
