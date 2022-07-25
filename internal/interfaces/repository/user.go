package repository

import (
	"context"

	"github.com/Ras96/go-clean-architecture-template/internal/domain/model"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent/user"
	"github.com/Ras96/go-clean-architecture-template/internal/usecases/repository"
)

type userRepositoryImpl struct {
	uc *ent.UserClient
}

func NewUserRepository(uc *ent.UserClient) repository.UserRepository {
	return &userRepositoryImpl{
		uc: uc,
	}
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id int) (model.User, error) {
	user, err := r.uc.Query().Where(user.IDEQ(id)).First(ctx)
	if err != nil {
		return model.User{}, convertError(err)
	}

	return model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (r *userRepositoryImpl) Create(ctx context.Context, params *repository.CreateUserParams) (model.User, error) {
	user, err := r.uc.Create().SetName(params.Name).SetEmail(params.Email).Save(ctx)
	if err != nil {
		return model.User{}, convertError(err)
	}

	return model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
