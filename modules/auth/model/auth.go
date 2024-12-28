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

type UserOutput struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	Status   string `json:"status"`
	RoleId   string `json:"role_id"`
	Email    string `json:"email"`
}
type UserResponse struct {
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	UserId   string `json:"user_id"`
	Status   string `json:"status"`
	RoleId   string `json:"role_id"`
	Email    string `json:"email"`
}
