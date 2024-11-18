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
	NameAccount string `json:"name_account" validate:"required"`
	Password    string `json:"password" validate:"required"`
}
