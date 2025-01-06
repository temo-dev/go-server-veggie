package common

import (
	"fmt"
	"strings"
)

// Custom Err
func ErrCannotCreateUser(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreateUser %s", entity))
}

func ErrCannotGetUser(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot get %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotGetUser %s", entity))
}

func ErrCannotDeleteUser(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot delete %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotDeleteUser %s", entity))
}

func ErrCannotUpdateUser(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotUpdateUser %s", entity))
}

func ErrUpdateUserIsNotChanged(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update %s", strings.ToLower(entity)), fmt.Sprintf("ErrUpdateUserIsNotChanged %s", entity))
}
