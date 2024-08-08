package events

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
)

type EventPublisher interface {
	Publish(topic string, event Event) error
}

type WatermillEventPublisher struct {
	publisher message.Publisher
}

func NewWatermillEventPublisher(publisher message.Publisher) *WatermillEventPublisher {
	return &WatermillEventPublisher{publisher: publisher}
}

func (p *WatermillEventPublisher) Publish(topic string, event Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	msg := message.NewMessage(event.ID, payload)
	return p.publisher.Publish(topic, msg)
}
