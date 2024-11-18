package model

import "errors"

// error
const (
	EntityName = "token"
)

var (
	ErrTokenMissing         = errors.New("wrong password")
	ErrInvalidTokenFormat   = errors.New("wrong format token")
	ErrAuthorizationMissing = errors.New("authorization header is missing")
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidClaims        = errors.New("invalid claims")
	ErrInvalidExpToken      = errors.New("token is expired")
	ErrInvalidAccount       = errors.New("invalid account")
)

//model
