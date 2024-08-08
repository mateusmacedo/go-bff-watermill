package events

type Event struct {
	ID    string
	Data  interface{}
	Event string
}
type EventHandler interface {
	Handle(event Event) error
	CanHandle(event Event) bool
}
