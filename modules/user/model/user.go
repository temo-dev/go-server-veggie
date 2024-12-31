package model

import (
	"errors"
	"time"
)

//error

const (
	EntityName = "User"
)

var (
	ErrorUserNotFound = errors.New("user not found")
	ErrorUpdateUser   = errors.New("update user failed")
	ErrorCreateUser   = errors.New("create user failed")
)

// Create User
type UserCreationType struct {
	UserName string `json:"user_name" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,alphanum"`
	Email    string `json:"email" validate:"required,email"`
}

// Get Users
type UserType struct {
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Update User
type UpdateUserType struct {
	UserId   string `json:"user_id" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
	Status   string `json:"status" validate:"required"`
}
