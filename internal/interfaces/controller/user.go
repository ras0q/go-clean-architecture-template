package controller

import (
	"context"
	"net/http"

	"github.com/Ras96/go-clean-architecture-template/internal/usecases/repository"
	"github.com/Ras96/go-clean-architecture-template/pkg/errors"
)

type UserController interface {
	GetUser(ctx context.Context, req *GetUserRequest) (GetUserResponse, int, error)
	PostUser(ctx context.Context, req *PostUserRequest) (PostUserResponse, int, error)
}

type userControllerImpl struct {
	ur repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) UserController {
	return &userControllerImpl{
		ur: userRepository,
	}
}

type (
	User struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	GetUserRequest struct {
		ID int `param:"id"`
	}

	GetUserResponse User

	PostUserRequest struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	PostUserResponse User
)

func (c *userControllerImpl) GetUser(ctx context.Context, req *GetUserRequest) (GetUserResponse, int, error) {
	user, err := c.ur.FindByID(ctx, req.ID)
	if err != nil {
		return GetUserResponse{}, errors.StatusCode(err), errors.Wrap(err, "userRepository.FindByID")
	}

	return GetUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, http.StatusOK, nil
}

func (c *userControllerImpl) PostUser(ctx context.Context, req *PostUserRequest) (PostUserResponse, int, error) {
	params := repository.CreateUserParams{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := c.ur.Create(ctx, &params)
	if err != nil {
		return PostUserResponse{}, errors.StatusCode(err), errors.Wrap(err, "userRepository.Create")
	}

	return PostUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, http.StatusCreated, nil
}
