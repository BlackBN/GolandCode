package main

/**
正向代理
*/

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Proxy struct {
}

func (*Proxy)ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Printf("Received request: %s %s %s\n",r.Method,r.Host,r.RemoteAddr)
	transport := http.DefaultTransport

	// 浅拷贝一个request 对象，避免后续修影响了源对象
	req := new(http.Request)
	*req = *r

	// 设置X-Forward-For 头部
	if clientIp,_,err := net.SplitHostPort(r.RemoteAddr);err ==nil{
		if prior,ok := req.Header["X-Forward-For"];ok{
			clientIp = strings.Join(prior,", ") + ", " + clientIp
		}
		req.Header.Set("X-Forward-For",clientIp)
	}

	// 构造新请求
	response,err:=transport.RoundTrip(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 获取响应数据并返回
	for k,v := range response.Header{
		for _,v1 := range v{
			w.Header().Add(k,v1)
		}
	}
	w.WriteHeader(response.StatusCode)
	io.Copy(w,response.Body)
	response.Body.Close()

}
// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	response, err := http.Get("http://127.0.0.1:9000/")
	if err != nil {
		// handle error
	}
	//程序在使用完回复后必须关闭回复的主体。
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func main() {
	//fmt.Println("Serve on :8001")
	//http.Handle("/",&Proxy{})
	//http.ListenAndServe("0.0.0.0:8001",nil)
	//create reverse proxy to listen gotty service
	//vmGottyUrl := "127.0.0.1:9000" //"192.168.156.20:8001"
	//proxyVm := NewMultipleHostsReverseProxy(&url.URL{
	//	Scheme: "http",
	//	Host:   vmGottyUrl,
	//})
	//log.Fatal(http.ListenAndServe(":8011", proxyVm))

	// 设置路由，如果访问/，则调用index方法
	//http.HandleFunc("/", ServeHTTP)

	// 启动web服务，监听9090端口
	//err := http.ListenAndServe(":8001", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	//http.HandleFunc("/echo", echo)
	//http.ListenAndServe("localhost:8080", nil)
	GetWebS()
}
//func echo(w http.ResponseWriter, r*http.Request) {
//	c, err := upgrader.Upgrade(w, r, nil)
//}



func GetWebS(){
	var origin = "http://127.0.0.1:8090/"
	var url = "ws://127.0.0.1:9000/ws"
	ws, err := websocket.Dial(url,"",origin)
	if err != nil {
		log.Fatal(err)
	}
	var message []byte
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	m, err := ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg[:m])

	ws.Close()//关闭连接
}

func NewMultipleHostsReverseProxy(target *url.URL) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = target.Path
			req.URL.User = target.User
		},
	}
}



// 代码运行之后，会在本地的 8080 端口启动代理服务。修改浏览器的代理为 127.0.0.1：:8080
// 再访问网站，可以验证代理正常工作，也能看到它在终端打印出所有的请求信息。