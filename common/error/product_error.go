package common

import (
	"fmt"
	"strings"
)

func ErrCannotCreateProduct(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create product %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreateProduct %s", entity))
}

func ErrCannotFindAllProduct(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find all product %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindAllProduct %s", entity))
}

func ErrCannotFindProductById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find product by id %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindProductById %s", entity))
}

func ErrCannotUpdateProduct(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update product %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotUpdateProduct %s", entity))
}
