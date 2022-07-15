package repository

import (
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
)

func convertError(err error) error {
	switch {
	case ent.IsNotFound(err):
		return errors.ErrNotFound
	default:
		return err
	}
}

// func convertErrorMaskNotFound(err error) error {
// 	return convertError(ent.MaskNotFound(err))
// }
