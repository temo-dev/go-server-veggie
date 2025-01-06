package common

import (
	"fmt"
	"strings"
)

func ErrCannotCreateCurrency(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreateCurrency %s", err))
}

func ErrCannotFindCurrencyById(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find %s by id", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindCurrencyById %s", err))
}
