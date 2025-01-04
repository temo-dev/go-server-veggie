package common

import (
	"fmt"
	"strings"
)

func ErrCannotCreateCurrency(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot create %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreateCurrency%s", err))
}
