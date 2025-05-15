package event

import (
	"context"

	domain "github.com/h22k/wordle-turkish-overengineering/server/internal/domain/event"
)

type HandlerFunc func(event domain.Event)

type listener interface {
	Listen(ctx context.Context, eventName string, handler HandlerFunc) error
}

type broker interface {
	Publish(event domain.Event)
}

type Dispatcher struct {
	listener listener
	broker   broker
}

func NewDispatcher(listener listener, broker broker) *Dispatcher {
	return &Dispatcher{
		listener: listener,
		broker:   broker,
	}
}

func (d *Dispatcher) Start(ctx context.Context, eventName string) error {
	return d.listener.Listen(ctx, eventName, d.broker.Publish)
}
