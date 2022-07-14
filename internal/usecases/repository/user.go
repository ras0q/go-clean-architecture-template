package repository

import (
	"context"

	"github.com/Ras96/go-clean-architecture-template/internal/domain/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id string) (model.User, error)
}
