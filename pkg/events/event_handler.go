package events

import "context"

type Event struct {
	ID   string
	Data interface{}
	Name string
}
type EventHandler interface {
	Handle(ctx context.Context, event Event) error
	CanHandle(event Event) bool
}
