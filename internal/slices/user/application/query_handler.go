package application

import "github.com/mateusmacedo/bff-watermill/internal/slices/user/domain"

type GetUserQuery struct {
	UserID string
}

type GetUserQueryHandler struct {
	userService domain.UserService
}

func NewGetUserQueryHandler(userService domain.UserService) *GetUserQueryHandler {
	return &GetUserQueryHandler{userService: userService}
}

func (h *GetUserQueryHandler) Handle(query GetUserQuery) (*domain.User, error) {
	return h.userService.GetUser(query.UserID)
}
