package pgsql

import (
	"context"
	"fmt"

	application "github.com/h22k/wordle-turkish-overengineering/server/internal/application/event"
	"github.com/h22k/wordle-turkish-overengineering/server/internal/domain/event"
	"github.com/jackc/pgx/v5"
)

type EventListener struct {
	conn *pgx.Conn
}

func NewEventListener(conn *pgx.Conn) *EventListener {
	return &EventListener{
		conn: conn,
	}
}

func (el *EventListener) Listen(ctx context.Context, eventName string, handler application.HandlerFunc) error {
	_, err := el.conn.Exec(ctx, "LISTEN "+eventName)

	if err != nil {
		return err
	}

	for {
		notification, err := el.conn.WaitForNotification(ctx)
		if err != nil {
			return err
		}
		fmt.Println(notification.Payload)
		handler(event.NewEvent(eventName, notification.Payload))
	}
}
