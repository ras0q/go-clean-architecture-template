package repository

import (
	"context"
	"fmt"

	"github.com/Ras96/go-clean-architecture-template/internal/domain/model"
	"github.com/Ras96/go-clean-architecture-template/internal/usecases/repository"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
)

type user struct {
	id    string
	name  string
	email string
}

type userRepositoryImpl struct {
	users map[string]user
}

func NewUserRepository() repository.UserRepository {
	return &userRepositoryImpl{
		users: map[string]user{
			"1": {"1", "Ras", "ras@example.com"},
			"2": {"2", "Cal", "cal@example.com"},
		},
	}
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id string) (model.User, error) {
	user, ok := r.users[id]
	if !ok {
		return model.User{}, errors.Wrap(errors.ErrNotFound, fmt.Sprintf("r.users[%s]", id))
	}

	return model.User{
		ID:    user.id,
		Name:  user.name,
		Email: user.email,
	}, nil
}
