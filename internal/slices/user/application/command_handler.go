package application

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/mateusmacedo/bff-watermill/internal/slices/user/domain"
	"github.com/mateusmacedo/bff-watermill/pkg/events"
)

type CreateUserCommand struct {
	Name  string
	Email string
}

type CreateUserCommandHandler struct {
	userService domain.UserService
	publisher   events.EventPublisher
}

func NewCreateUserCommandHandler(userService domain.UserService, publisher events.EventPublisher) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{userService: userService, publisher: publisher}
}

func (h *CreateUserCommandHandler) Handle(cmd CreateUserCommand) (string, error) {
	user := &domain.User{Name: cmd.Name, Email: cmd.Email}

	event, err := h.userService.CreateUser(user)
	if err != nil {
		return "", err
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	msg := message.NewMessage(user.ID, payload)
	msg.Metadata.Set("type", "user.created")
	err = h.publisher.Publish("user_events", msg)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
