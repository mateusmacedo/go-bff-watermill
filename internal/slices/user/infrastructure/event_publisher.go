package infrastructure

import (
	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/mateusmacedo/bff-watermill/pkg/events"
)

type WatermillEventPublisher struct {
	publisher message.Publisher
}

func NewWatermillEventPublisher(publisher message.Publisher) *WatermillEventPublisher {
	return &WatermillEventPublisher{publisher: publisher}
}

func (p *WatermillEventPublisher) Publish(topic string, event events.Event) error {
	msg := message.NewMessage(event.ID, []byte(event.Data.(string)))
	return p.publisher.Publish(topic, msg)
}
