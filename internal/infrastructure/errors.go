package infrastructure

import "github.com/Ras96/go-clean-architecture-template/pkg/errors"

type errorCode uint

var _ errors.CodeIsImplementedBy[errorCode]

const (
	ECEchoError errorCode = iota
	ECEntError
)

func (c errorCode) Messsage() string {
	switch c {
	case ECEchoError:
		return "echo error"
	case ECEntError:
		return "ent error"
	default:
		return "unknown error"
	}
}
