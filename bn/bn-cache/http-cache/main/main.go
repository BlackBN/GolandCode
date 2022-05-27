package main

import (
	cache "GolandCode/bn/bn-cache/http-cache"
	"fmt"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	cache.NewGroup("scores", cache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}), 2<<10)
	selfAddress := "127.0.0.1:7869"
	httppool := cache.NewHttpPool(selfAddress)
	log.Printf("cache is running at %s\n", selfAddress)
	log.Fatal(http.ListenAndServe(selfAddress, httppool))
}
