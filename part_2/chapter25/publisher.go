package chapter25

import (
	"context"
	"sync"
)

type Publisher struct {
	ctx          context.Context
	subscriberCh chan chan<- string
	publishCh    chan string
	subscribers  []chan<- string
}

var Wg sync.WaitGroup

func NewPublisher(ctx context.Context) *Publisher {
	return &Publisher{
		ctx:          ctx,
		subscriberCh: make(chan chan<- string),
		publishCh:    make(chan string),
		subscribers:  make([]chan<- string, 0),
	}
}

func (p *Publisher) Subscribe(sub chan<- string) {
	p.subscriberCh <- sub
}

func (p *Publisher) Publish(msg string) {
	p.publishCh <- msg
}

func (p *Publisher) Update() {
	for {
		select {
		case sub := <-p.subscriberCh:
			p.subscribers = append(p.subscribers, sub)
		case msg := <-p.publishCh:
			for _, v := range p.subscribers {
				v <- msg
			}
		case <-p.ctx.Done():
			Wg.Done()
			return
		}
	}
}
