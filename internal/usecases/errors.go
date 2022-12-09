package usecases

import "github.com/ras0q/go-clean-architecture-template/pkg/errors"

type ErrorCode uint

var _ errors.CodeIsImplementedBy[ErrorCode]

const (
	ECNotFound ErrorCode = iota
	ECInvalidArgument
	ECAlreadyExists

	ECInternal
)

var errorMessages = map[ErrorCode]string{
	ECNotFound:        "not found",
	ECInvalidArgument: "invalid argument",
	ECAlreadyExists:   "already exists",
	ECInternal:        "internal error",
}

func (c ErrorCode) Messsage() string {
	if m, ok := errorMessages[c]; ok {
		return m
	}

	return "unknown error"
}
