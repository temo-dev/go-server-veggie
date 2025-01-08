package common

import (
	"fmt"
	"strings"
)

func ErrCannotFindBrand(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot find %s ", strings.ToLower(entity)), fmt.Sprintf("ErrCannotFindBrand %s", err))
}

func ErrCannotUpdateBrand(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot update %s ", strings.ToLower(entity)), fmt.Sprintf("ErrCannotUpdateBrand %s", err))
}
