package controller

import (
	"context"

	"github.com/Ras96/go-clean-architecture-template/internal/usecases/repository"
)

type UserController interface {
	GetUser(ctx context.Context, req *GetUserRequest) (GetUserResponse, error)
	PostUser(ctx context.Context, req *PostUserRequest) (PostUserResponse, error)
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

func (c *userControllerImpl) GetUser(ctx context.Context, req *GetUserRequest) (GetUserResponse, error) {
	user, err := c.ur.FindByID(ctx, req.ID)
	if err != nil {
		return GetUserResponse{}, newHTTPError(err)
	}

	return GetUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (c *userControllerImpl) PostUser(ctx context.Context, req *PostUserRequest) (PostUserResponse, error) {
	params := repository.CreateUserParams{
		Name:  req.Name,
		Email: req.Email,
	}

	user, err := c.ur.Create(ctx, &params)
	if err != nil {
		return PostUserResponse{}, newHTTPError(err)
	}

	return PostUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
