package application

import (
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
	err := h.userService.CreateUser(user)
	if err != nil {
		return "", err
	}

	event := events.Event{
		ID:      user.ID,
		Payload: cmd,
		Name:    "UserCreated",
	}

	err = h.publisher.Publish("user_events", event)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}
