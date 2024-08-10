package events

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type EventPublisher interface {
	Publish(topic string, event *message.Message) error
}

type WatermillEventPublisher struct {
	publisher message.Publisher
}

func NewWatermillEventPublisher(publisher message.Publisher) *WatermillEventPublisher {
	return &WatermillEventPublisher{publisher: publisher}
}

func (p *WatermillEventPublisher) Publish(topic string, event *message.Message) error {
	return p.publisher.Publish(topic, event)
}
