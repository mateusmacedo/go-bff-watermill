package application

import (
	"context"

	_application "github.com/mateusmacedo/bff-watermill/pkg/application"
	"github.com/mateusmacedo/bff-watermill/pkg/events"
)

type UserCreatedHandler struct {
	logger _application.AppLogger
}

func NewUserCreatedHandler(logger _application.AppLogger) events.EventHandler {
	return &UserCreatedHandler{logger: logger}
}

func (h *UserCreatedHandler) Handle(ctx context.Context, event events.Event) error {
	h.logger.Info(ctx, "Evento recebido", map[string]interface{}{
		"event": event,
	})

	return nil
}

func (h *UserCreatedHandler) CanHandle(event events.Event) bool {
	return event.Name == "UserCreated"
}
