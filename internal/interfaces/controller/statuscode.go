package controller

import (
	"net/http"

	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
)

func statusCode(err error) int {
	switch {
	case errors.Is(err, errors.ErrValidate):
		return http.StatusBadRequest
	case errors.Is(err, errors.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, errors.ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
