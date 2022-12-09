//go:generate go run github.com/golang/mock/mockgen@latest -source=$GOFILE -destination=mock_$GOPACKAGE/mock_$GOFILE

package repository

import (
	"context"

	"github.com/ras0q/go-clean-architecture-template/internal/domain"
)

type (
	UserRepository interface {
		FindByID(ctx context.Context, id int) (domain.User, error)
		Create(ctx context.Context, params *CreateUserParams) (domain.User, error)
	}

	CreateUserParams struct {
		Name  string
		Email string
	}
)
