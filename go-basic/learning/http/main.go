package main

import "net/http"

func main() {

	//127.0.0.1:4378/assert/
	http.ListenAndServe(":4378", http.StripPrefix("/assert", http.FileServer(http.Dir("/Users/bn/GoProject/src/GolandCode"))))
}
