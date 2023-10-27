package main

import (
	"GolandCode/design/action/pub-sub/pubsub"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	p := pubsub.NewPublisher(100, 1000)
	defer p.Close()
	all := p.Subscribe()

	test := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok && strings.Contains(s, "xujiayu") {
			return true
		}
		return false
	})

	p.Publish("xujiayu test")
	p.Publish("panshuyi test")
	go func() {
		for v := range all {
			fmt.Println("other test : ", v)
		}
	}()

	go func() {
		for v := range test {
			fmt.Println("get xujiayu test : ", v)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
