package domain

import "errors"

type User struct {
	ID    string
	Name  string
	Email string
}

var ErrUserNotFound = errors.New("user not found")
