package event

import (
	"sync"

	"github.com/h22k/wordle-turkish-overengineering/server/internal/domain/event"
)

type Broker struct {
	subscribers map[chan event.Event]struct{}
	mu          sync.Mutex
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[chan event.Event]struct{}),
	}
}

func (b *Broker) Subscribe() chan event.Event {
	subscriber := make(chan event.Event)
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers[subscriber] = struct{}{}
	return subscriber
}

func (b *Broker) Unsubscribe(subscriber chan event.Event) {
	b.mu.Lock()
	defer b.mu.Unlock()
	delete(b.subscribers, subscriber)
	close(subscriber)
}

func (b *Broker) Publish(event event.Event) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for subscriber := range b.subscribers {
		select {
		case subscriber <- event:
		default:
		}
	}
}
