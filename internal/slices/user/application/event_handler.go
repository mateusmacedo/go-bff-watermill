package application

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/mateusmacedo/bff-watermill/internal/slices/user/domain"
	_application "github.com/mateusmacedo/bff-watermill/pkg/application"
	"github.com/mateusmacedo/bff-watermill/pkg/events"
)

type UserCreatedHandler struct {
	logger _application.AppLogger
}

func NewUserCreatedHandler(logger _application.AppLogger) events.EventHandler {
	return &UserCreatedHandler{logger: logger}
}

func (h *UserCreatedHandler) Handle(ctx context.Context, msg *message.Message) error {
	var event domain.UserCreatedEvent
	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		return err
	}
	h.logger.Info(ctx, "Evento recebido", map[string]interface{}{
		"event": event,
	})

	return nil
}

func (h *UserCreatedHandler) CanHandle(msg *message.Message) bool {
	return msg.Metadata.Get("type") == "user.created"
}
