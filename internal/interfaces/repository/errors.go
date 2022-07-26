//nolint:nlreturn
package repository

import (
	"log"

	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
)

func convertError(err error) error {
	switch {
	case err == nil:
		return nil
	case ent.IsNotFound(err):
		// not logging here because it's a common error
		return errors.ErrNotFound
	case ent.IsValidationError(err):
		log.Printf("[ERROR]: %v\n", err)
		return errors.ErrValidate
	case ent.IsConstraintError(err):
		log.Printf("[ERROR]: %v\n", err)
		return errors.ErrConflict
	default:
		log.Printf("[ERROR]: %v\n", err)
		return err
	}
}

// func convertErrorMaskNotFound(err error) error {
// 	return convertError(ent.MaskNotFound(err))
// }
