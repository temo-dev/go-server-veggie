package model

import (
	"errors"
)

//error

const (
	EntityName = "User"
)

var (
	ErrorUserNotFound           = errors.New("user not found")
	ErrorUpdateUser             = errors.New("update user failed")
	ErrorCreateUser             = errors.New("create user failed")
	ErrorUpdateUserIsNotChanged = errors.New("update user is not changed")
)

// Create User
type UserCreationType struct {
	UserName string `json:"user_name" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,alphanum"`
	Email    string `json:"email" validate:"required,email"`
}

// Get Users
type UserType struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Status   string `json:"status"`
}
