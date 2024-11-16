package common

import (
	"fmt"
	"strings"
)

// auth Error
func ErrCannotLogin(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("cannot login %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotLogin %s", entity))
}
