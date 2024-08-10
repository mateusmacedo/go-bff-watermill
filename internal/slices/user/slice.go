package user

import (
	"github.com/go-chi/chi/v5"

	"github.com/mateusmacedo/bff-watermill/internal/slices/user/application"
	"github.com/mateusmacedo/bff-watermill/internal/slices/user/domain"
	"github.com/mateusmacedo/bff-watermill/internal/slices/user/infrastructure"
	iface "github.com/mateusmacedo/bff-watermill/internal/slices/user/interface"
	"github.com/mateusmacedo/bff-watermill/pkg/events" // Certifique-se de que o pacote correto está sendo importado
)

type UserSlice struct {
	httpHandler *iface.UserHTTPHandler
}

func NewUserSlice(publisher events.EventPublisher, subscriber events.EventSubscriber) *UserSlice {
	userRepo := infrastructure.NewUserRepository()
	userService := domain.NewUserService(userRepo)

	commandHandler := application.NewCreateUserCommandHandler(userService, publisher)
	queryHandler := application.NewGetUserQueryHandler(userService)

	httpHandler := iface.NewUserHTTPHandler(commandHandler, queryHandler)

	return &UserSlice{
		httpHandler: httpHandler,
	}
}

func (s *UserSlice) RegisterRoutes(router *chi.Mux) {
	s.httpHandler.RegisterRoutes(router)
}
