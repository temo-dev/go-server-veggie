package common

import (
	"fmt"
	"strings"
)

func ErrCannotCreateAPP(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create attitude product package %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreateAPP %s", entity))
}

func ErrCannotFindAPP(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find attitude product package %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindAPP %s", entity))
}

func ErrCannotFindAPPById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find attitude product package by id %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindAPPById %s", entity))
}

func ErrCannotUpdateAPP(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update attitude product package %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotUpdateAPP %s", entity))
}
