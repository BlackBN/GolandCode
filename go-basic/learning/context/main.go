package main

import (
	"context"
	"fmt"
	"sync"
)

func callRpc(ctx context.Context, url string, isSuccess bool) error {
	result := make(chan int)
	err := make(chan error)

	go func() {
		if isSuccess {
			result <- 1
		} else {
			err <- fmt.Errorf("call %s is error\n", url)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("call %s has done\n", url)
		return ctx.Err()
	case e := <-err:
		return e
	case <-result:
		return nil

	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	err := callRpc(ctx, "http://aaaaa", true)
	if err != nil {
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := callRpc(ctx, "http://bbbbb", false)
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := callRpc(ctx, "http://ccccc", true)
		if err != nil {
			cancel()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := callRpc(ctx, "http://ddddd", true)
		if err != nil {
			cancel()
		}
	}()

	wg.Wait()

}
