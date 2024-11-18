package common

import (
	"fmt"
	"strings"
)

// auth Error
func ErrCannotLogin(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot login %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotLogin %s", entity))
}

// token Error
func ErrTokenMissing(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("token is missing %s", strings.ToLower(entity)), fmt.Sprintf("ErrTokenMissing %s", entity))
}

func ErrInvalidTokenFormat(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("wrong format token %s", strings.ToLower(entity)), fmt.Sprintf("ErrInvalidTokenFormat %s", entity))
}

func ErrAuthorizationMissing(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("authorization header is missing %s", strings.ToLower(entity)), fmt.Sprintf("ErrAuthorizationMissing %s", entity))
}

func ErrInvalidToken(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("invalid token %s", strings.ToLower(entity)), fmt.Sprintf("ErrInvalidToken %s", entity))
}

func ErrInvalidClaims(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("invalid claims %s", strings.ToLower(entity)), fmt.Sprintf("ErrInvalidClaims %s", entity))
}

func ErrInvalidExpToken(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("token is expired %s", strings.ToLower(entity)), fmt.Sprintf("ErrInvalidExpToken %s", entity))
}

func ErrInvalidAccount(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("invalid account %s", strings.ToLower(entity)), fmt.Sprintf("ErrInvalidAccount %s", entity))
}
