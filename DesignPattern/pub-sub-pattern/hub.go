package main

import (
	"context"
	"fmt"
	"sync"
)

type hub struct {
	sync.Mutex
	name string
	subs map[*subscriber]struct{}
}

func (h *hub) subscribe(ctx context.Context, sub *subscriber) error {
	h.Lock()
	_, exist := h.subs[sub]
	if exist {
		return fmt.Errorf("subscribe already exist %s", sub.String())
	}
	h.subs[sub] = struct{}{}
	h.Unlock()

	go func() {
		for {
			select {
			case <-sub.quick:
				return
			case <-ctx.Done():
				h.Lock()
				delete(h.subs, sub)
				h.Unlock()
				return
			}
		}
	}()

	go sub.run(ctx)

	return nil
}

func (h *hub) unsubscribe(ctx context.Context, sub *subscriber) error {
	h.Lock()
	_, exist := h.subs[sub]
	if !exist {
		return fmt.Errorf("subscribe not exist %s", sub.String())
	}
	delete(h.subs, sub)
	h.Unlock()

	close(sub.quick)

	return nil
}

func (h *hub) publish(ctx context.Context, msg *Message) {
	h.Lock()
	for sub := range h.subs {
		sub.publish(ctx, msg)
	}
	h.Unlock()
}

func NewHub(name string) *hub {
	return &hub{
		name: name,
		subs: make(map[*subscriber]struct{}),
	}
}
