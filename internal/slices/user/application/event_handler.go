package application

import (
	"encoding/json"

	"github.com/mateusmacedo/bff-watermill/pkg/events"
	"github.com/mateusmacedo/bff-watermill/pkg/infrastructure"
)

// UserCreatedHandler processa eventos de usuário criado
type UserCreatedHandler struct {
	logger *infrastructure.Logger
}

func NewUserCreatedHandler(logger *infrastructure.Logger) events.EventHandler {
	return &UserCreatedHandler{logger: logger}
}

func (h *UserCreatedHandler) Handle(event events.Event) error {
	eventMarshaled, err := json.Marshal(event)
	if err != nil {
		h.logger.Error("Erro ao desserializar evento" + err.Error())
		return err
	}

	h.logger.Info(string(eventMarshaled))

	return nil
}

func (h *UserCreatedHandler) CanHandle(event events.Event) bool {
	return event.Event == "UserCreated"
}
