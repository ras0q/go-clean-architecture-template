//nolint:nlreturn
package repository

import (
	"github.com/ras0q/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/ras0q/go-clean-architecture-template/internal/usecases"
	"github.com/ras0q/go-clean-architecture-template/pkg/errors"
)

func convertError(err error) error {
	if err == nil {
		return nil
	}

	var code usecases.ErrorCode
	switch {
	case ent.IsNotFound(err):
		code = usecases.ECNotFound
	case ent.IsValidationError(err):
		code = usecases.ECInvalidArgument
	case ent.IsConstraintError(err):
		code = usecases.ECAlreadyExists
	default:
		return err
	}

	return errors.Wrap(code, err)
}
