package errors

import (
	"errors"
	"fmt"
	"net/http"
)

// Rule:
// 1. Define basic errors in this package
// 2. Use `errors.Wrap` when returning error from a function
// 3. Use `errors.Is` or `err != nil` to check error

var (
	ErrNotFound = errors.New("not found")
	ErrBind     = errors.New("bind error")
	ErrValidate = errors.New("validate error")
	ErrConflict = errors.New("invalid argument")
	ErrInternal = errors.New("internal error")
)

// StatusCode returns the appropriate status code for the error content.
func StatusCode(err error) int {
	switch {
	case Is(err, ErrValidate):
		return http.StatusBadRequest
	case Is(err, ErrNotFound):
		return http.StatusNotFound
	case Is(err, ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// Wrap wraps standard fmt.Errorf
func Wrap(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// Is wraps standard errors.Is
func Is(err error, target error) bool {
	return errors.Is(err, target)
}

// As wraps standard errors.As
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
