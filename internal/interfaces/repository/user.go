package repository

import (
	"context"

	"github.com/ras0q/go-clean-architecture-template/internal/domain"
	"github.com/ras0q/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/ras0q/go-clean-architecture-template/internal/interfaces/repository/ent/user"
	"github.com/ras0q/go-clean-architecture-template/internal/usecases/repository"
)

type userRepositoryImpl struct {
	uc *ent.UserClient
}

func NewUserRepository(uc *ent.UserClient) repository.UserRepository {
	return &userRepositoryImpl{
		uc: uc,
	}
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id int) (domain.User, error) {
	user, err := r.uc.Query().Where(user.IDEQ(id)).First(ctx)
	if err != nil {
		return domain.User{}, convertError(err)
	}

	return domain.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, params *repository.CreateUserParams) (domain.User, error) {
	user, err := r.uc.Create().SetName(params.Name).SetEmail(params.Email).Save(ctx)
	if err != nil {
		return domain.User{}, convertError(err)
	}

	return domain.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
