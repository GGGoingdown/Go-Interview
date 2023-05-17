package main

import (
	"context"
	"fmt"
	"log"
)

type Message struct {
	data string
}

type subscriber struct {
	name    string
	handler chan *Message
	quick   chan struct{}
}

func (s *subscriber) String() string {
	return fmt.Sprintf("SUB %s", s.name)
}

func (s *subscriber) run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("SUB[%s]::context done", s.name)
			return
		case msg := <-s.handler:
			log.Printf("SUB[%s]::receive message: %s\n", s.name, msg.data)
		case <-s.quick:
			log.Printf("SUB[%s]::unsubscribe", s.name)
			return
		}
	}
}

func (s *subscriber) publish(ctx context.Context, msg *Message) {
	select {
	case <-ctx.Done():
		log.Printf("SUB[%s]::context done", s.name)
		return
	case s.handler <- msg:
	}
}

func NewSubsriber(name string) *subscriber {
	return &subscriber{
		name:    name,
		handler: make(chan *Message, 1),
		quick:   make(chan struct{}, 1),
	}
}
