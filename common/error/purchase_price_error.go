package common

import (
	"fmt"
	"strings"
)

func ErrCannotCreatePurchasePrice(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create purchase price %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreatePurchasePrice %s", entity))
}

func ErrCannotFindAllPurchasePrice(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find all purchase price %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindAllPurchasePrice %s", entity))
}

func ErrCannotFindPurchasePriceById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find purchase price by id %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindPurchasePriceById %s", entity))
}

func ErrCannotUpdatePurchasePrice(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update purchase price %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotUpdatePurchasePrice %s", entity))
}

func ErrCannotDeletePurchasePrice(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot delete purchase price %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotDeletePurchasePrice %s", entity))
}
