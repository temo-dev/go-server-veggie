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
)

// Create User
type UserCreationType struct {
	NameAccount string `json:"name_account" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Role        string `json:"role" validate:"required"`
}

// Get Users
type UserType struct {
	Id          string     `json:"id"`
	NameAccount *string    `json:"name_account"`
	Role        *string    `json:"role"`
	Status      *string    `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// Update User
type UpdateUserType struct {
	Id          string `json:"id" validate:"required"`
	NameAccount string `json:"name_account" validate:"required"`
	Status      string `json:"status" validate:"required"`
	Role        string `json:"role" validate:"required"`
}
