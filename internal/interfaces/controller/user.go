package controller

import (
	"context"

	"github.com/Ras96/go-clean-architecture-template/internal/usecases/repository"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
)

type UserController interface {
	GetUser(ctx context.Context, req *GetUserRequest) (GetUserResponse, error)
}

type userControllerImpl struct {
	ur repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) UserController {
	return &userControllerImpl{
		ur: userRepository,
	}
}

type GetUserRequest struct {
	ID string `param:"id"`
}

type GetUserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *userControllerImpl) GetUser(ctx context.Context, req *GetUserRequest) (GetUserResponse, error) {
	user, err := c.ur.FindByID(ctx, req.ID)
	if err != nil {
		return GetUserResponse{}, errors.Wrap(err, "userRepository.FindByID")
	}

	return GetUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
