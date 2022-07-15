package repository

import (
	"context"

	"github.com/Ras96/go-clean-architecture-template/internal/domain/model"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent"
	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/repository/ent/user"
	"github.com/Ras96/go-clean-architecture-template/internal/usecases/repository"
)

type userRepositoryImpl struct {
	*ent.UserClient
}

func NewUserRepository(uc *ent.UserClient) repository.UserRepository {
	return &userRepositoryImpl{uc}
}

func (r *userRepositoryImpl) FindByID(ctx context.Context, id int) (model.User, error) {
	user, err := r.Query().Where(user.IDEQ(id)).First(ctx)
	if cerr := convertError(err); cerr != nil {
		return model.User{}, cerr
	}

	return model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
