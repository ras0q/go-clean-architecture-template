package controller

import (
	"net/http"

	"github.com/Ras96/go-clean-architecture-template/internal/usecases"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
)

type StatusCode uint

var _ errors.CodeIsImplementedBy[StatusCode]

func (c StatusCode) Messsage() string {
	return http.StatusText(int(c))
}

func newHTTPError(err error) error {
	var code StatusCode

	if ce, ok := err.(*errors.CodeError[usecases.ErrorCode]); ok {
		switch ce.Code {
		case usecases.ECNotFound:
			code = http.StatusNotFound
		case usecases.ECInvalidArgument:
			code = http.StatusBadRequest
		case usecases.ECAlreadyExists:
			code = http.StatusConflict
		case usecases.ECInternal:
			code = http.StatusInternalServerError
		}
	}

	return errors.Wrap(code, err)
}
