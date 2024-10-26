package use_cases

import (
	"context"
	"go-auth-api/internal/application/validators"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type createUserUseCaseImpl struct {
	userService adapters.UserService
}

func CreateUserUseCaseConstructor(userService adapters.UserService) adapters.CreateUserUseCase {
	return &createUserUseCaseImpl{
		userService: userService,
	}
}

func (c createUserUseCaseImpl) Execute(ctx context.Context, user *dtos.UsersDto) (*types.User, exceptions.ErrorType) {
	validationErr := validators.Validate(*user, ctx)

	if validationErr != nil {
		return nil, validationErr
	}

	createdUser, err := c.userService.NewUser(ctx, user)

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
