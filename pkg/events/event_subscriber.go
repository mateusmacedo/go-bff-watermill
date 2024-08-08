package events

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type EventSubscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
}

type WatermillEventSubscriber struct {
	subscriber message.Subscriber
}

func NewWatermillEventSubscriber(subscriber message.Subscriber) *WatermillEventSubscriber {
	return &WatermillEventSubscriber{subscriber: subscriber}
}

func (s *WatermillEventSubscriber) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	return s.subscriber.Subscribe(ctx, topic)
}
