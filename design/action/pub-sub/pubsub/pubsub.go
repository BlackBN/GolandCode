package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)

type Publisher struct {
	m           sync.Mutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPublisher(buffer int, timeout time.Duration) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     timeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	defer p.m.Unlock()
	p.subscribers[ch] = topic

	return ch
}

func (p *Publisher) Subscribe() chan interface{} {
	return p.SubscribeTopic(nil)
}

func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subscribers, sub)
	close(sub)
}

func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.SendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) SendTopic(sub chan interface{}, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}
