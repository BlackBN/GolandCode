package singleflight

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

//用阻塞式读防止缓存击穿
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	//当有大量请求同时访问同一个key，可以使用waitGroup阻塞住除第一个请求以外的请求。
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		//fmt.Printf("%d ,before wait %v,%v\n", GoID(), c.val, c.err)
		c.wg.Wait()
		//fmt.Printf("%d after wait %v,%v\n", GoID(), c.val, c.err)
		return c.val, c.err
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()
	//fmt.Printf("%d before execute fn\n", GoID())
	//第一个请求执行函数
	c.val, c.err = fn()
	//fmt.Printf("%d after execute fn\n", GoID())
	c.wg.Done()
	//fmt.Printf("%d after wg done fn\n", GoID())

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()
	return c.val, c.err
}

//获取到 goroutine id，即线程id
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	//fmt.Printf("stack : %s \n", string(buf[:n]))
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
