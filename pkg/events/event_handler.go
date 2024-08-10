package events

import "context"

type Event struct {
	ID      string
	Name    string
	Payload interface{}
}

type EventHandler interface {
	Handle(ctx context.Context, event Event) error
	CanHandle(event Event) bool
}
