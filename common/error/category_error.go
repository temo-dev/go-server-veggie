package common

import (
	"fmt"
	"strings"
)

func ErrCannotCreateCategory(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreateCategory %s", entity))
}

func ErrCannotFindCategories(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindCategories %s", entity))
}

func ErrCannotFindCategoryById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find category by id %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindCategoryById %s", entity))
}

func ErrCannotDeleteCategoryById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot delete category by id %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotDeleteCategoryById %s", entity))
}

func ErrUpdateCategory(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update category %s", strings.ToLower(entity)), fmt.Sprintf("ErrUpdateCategory %s", entity))
}

func ErrUpdateCategoryIsNotChanged(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update category which is not changed %s", strings.ToLower(entity)), fmt.Sprintf("ErrUpdateCategoryIsNotChanged %s", entity))
}
