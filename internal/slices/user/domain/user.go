package domain

import "errors"

type User struct {
	ID    string
	Name  string
	Email string
}

type UserCreatedEvent struct {
	Name    string
	Payload struct {
		Name  string
		Email string
	}
}

var ErrUserNotFound = errors.New("user not found")
