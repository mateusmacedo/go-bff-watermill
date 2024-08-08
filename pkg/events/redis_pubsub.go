package events

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/redis/go-redis/v9"
)

// NewWatermillRedisPublisher cria um publisher usando Redis Streams
func NewWatermillRedisPublisher(redisClient redis.UniversalClient, logger watermill.LoggerAdapter) *WatermillEventPublisher {
	publisher, err := redisstream.NewPublisher(redisstream.PublisherConfig{
		Client: redisClient,
	}, logger)
	if err != nil {
		log.Fatalf("Erro ao criar publisher: %v", err)
	}
	return NewWatermillEventPublisher(publisher)
}

// NewWatermillRedisSubscriber cria um subscriber usando Redis Streams
func NewWatermillRedisSubscriber(redisClient redis.UniversalClient, logger watermill.LoggerAdapter) *WatermillEventSubscriber {
	subscriber, err := redisstream.NewSubscriber(redisstream.SubscriberConfig{
		Client:        redisClient,
		ConsumerGroup: "my_group",
		Consumer:      "my_consumer",
	}, logger)
	if err != nil {
		log.Fatalf("Erro ao criar subscriber: %v", err)
	}
	return NewWatermillEventSubscriber(subscriber)
}
