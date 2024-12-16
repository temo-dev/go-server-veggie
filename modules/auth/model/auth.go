package model

import "errors"

// error
const (
	EntityName = "auth"
)

var (
	ErrWrongPassword = errors.New("wrong password or account")
)

// model
type LoginInput struct {
	UserName string `json:"user_name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LogoutInput struct {
	Token string `json:"token" validate:"required"`
}
