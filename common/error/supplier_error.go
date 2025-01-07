package common

import (
	"fmt"
	"strings"
)

func ErrCannotCreateSupplier(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreateSupplier %s", err))
}

func ErrCannotFindSupplierById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find %s by id", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindSupplierById %s", err))
}

func ErrCannotUpdateSupplierIsNotChanged(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update %s is not changed", strings.ToLower(entity)), fmt.Sprintf("ErrCannotUpdateSupplierIsNotChanged %s", err))
}
