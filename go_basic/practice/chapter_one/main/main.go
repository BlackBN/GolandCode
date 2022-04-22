package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)
func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("ch insert one value")
		ch <- 1
		fmt.Println("test dd ")
	}()

	go func() {
		fmt.Println("ch read one value")
		<- ch
		fmt.Println("ch read one value is success")
		//time.Sleep(1 * time.Second)


	}()
	fmt.Println("end")
	//go produce(ch)
	//go consumer(ch)
	time.Sleep(1 * time.Second)

	//create reverse proxy to listen gotty service
	//vmGottyUrl := "127.0.0.1:9195"
	//proxyVm := NewMultipleHostsReverseProxy([]*url.URL{
	//	{
	//		Scheme: "http",
	//		Host:   vmGottyUrl,
	//	},
	//})
	//log.Fatal(http.ListenAndServe("9087", proxyVm))
}

func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}
