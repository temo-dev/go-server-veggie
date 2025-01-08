package common

import (
	"fmt"
	"strings"
)

func ErrCannotCreateSubCategory(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreateSubCategory %s", entity))
}

func ErrCannotFindSubCategories(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindSubCategories %s", entity))
}

func ErrCannotFindSubCategoryById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find category by id %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindSubCategoryById %s", entity))
}

func ErrCannotDeleteSubCategoryById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot delete category by id %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotDeleteSubCategoryById %s", entity))
}

func ErrUpdateSubCategory(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update category %s", strings.ToLower(entity)), fmt.Sprintf("ErrUpdateSubCategory %s", entity))
}

func ErrUpdateSubCategoryIsNotChanged(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update category which is not changed %s", strings.ToLower(entity)), fmt.Sprintf("ErrUpdateSubCategoryIsNotChanged %s", entity))
}
