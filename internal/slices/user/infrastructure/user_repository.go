package infrastructure

import "github.com/mateusmacedo/bff-watermill/internal/slices/user/domain"

type UserRepository struct {
	users map[string]*domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]*domain.User)}
}

func (r *UserRepository) Save(user *domain.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	user, exists := r.users[id]
	if !exists {
		return &domain.User{}, domain.ErrUserNotFound
	}
	return user, nil
}
