package events

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type EventHandler interface {
	Handle(ctx context.Context, event *message.Message) error
	CanHandle(event *message.Message) bool
}
