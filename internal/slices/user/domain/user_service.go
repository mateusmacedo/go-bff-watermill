package domain

import (
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(user *User) (UserCreatedEvent, error)
	GetUser(id string) (*User, error)
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

type userService struct {
	repo UserRepository
}

func (s *userService) CreateUser(user *User) (UserCreatedEvent, error) {
	user.ID = generateUserID() // Gerando um ID único dinâmico
	err := s.repo.Save(user)
	if err != nil {
		return UserCreatedEvent{}, err
	}

	return UserCreatedEvent{Name: "user.created", Payload: struct {
		Name  string
		Email string
	}{Name: user.Name, Email: user.Email}}, nil
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
