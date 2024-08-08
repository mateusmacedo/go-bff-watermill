package infrastructure

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/mateusmacedo/bff-watermill/pkg/events"
)

type WatermillEventSubscriber struct {
	subscriber message.Subscriber
}

func NewWatermillEventSubscriber(subscriber message.Subscriber) *WatermillEventSubscriber {
	return &WatermillEventSubscriber{subscriber: subscriber}
}

func (s *WatermillEventSubscriber) Subscribe(topic string) (<-chan events.Event, error) {
	ctx := context.Background() // Adicionando o contexto padrão
	messages, err := s.subscriber.Subscribe(ctx, topic)
	if err != nil {
		return nil, err
	}

	eventsChan := make(chan events.Event)
	go func() {
		defer close(eventsChan)
		for msg := range messages {
			eventsChan <- events.Event{
				ID:    msg.UUID,
				Data:  string(msg.Payload),
				Event: msg.Metadata.Get("event"),
			}
			msg.Ack() // Confirmar o processamento da mensagem
		}
	}()

	return eventsChan, nil
}
