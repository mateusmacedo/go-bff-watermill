package domain

import (
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *User) error
	GetUser(id string) (*User, error)
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

type userService struct {
	repo UserRepository
}

func (s *userService) CreateUser(user *User) error {
	user.ID = generateUserID() // Gerando um ID único dinâmico
	return s.repo.Save(user)
}

func (s *userService) GetUser(id string) (*User, error) {
	return s.repo.FindByID(id)
}

type UserRepository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
}

// Função para gerar um ID único
func generateUserID() string {
	return uuid.New().String()
}
