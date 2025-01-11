package common

import (
	"fmt"
	"strings"
)

func ErrCannotCreateTag(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create tag %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreateTag %s", entity))
}

func ErrCannotFindAllTag(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find all tag %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindAllTag %s", entity))
}

func ErrCannotFindTagById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find tag by id %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindTagById %s", entity))
}

func ErrCannotUpdateTag(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update tag %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotUpdateTag %s", entity))
}

func ErrCannotDeleteTag(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot delete tag %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotDeleteTag %s", entity))
}
