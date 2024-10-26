package use_cases

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/application/validators"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type updateUserUseCaseImpl struct {
	userService adapters.UserService
}

func UpdateUserUseCaseConstructor(userService adapters.UserService) adapters.UpdateUserUseCaseAdapter {
	return &updateUserUseCaseImpl{
		userService: userService,
	}
}

func (u updateUserUseCaseImpl) Execute(ctx context.Context, data *dtos.UpdateUserDto, id uuid.UUID) (*types.User, exceptions.ErrorType) {
	var err exceptions.ErrorType

	err = validators.Validate(*data, ctx)

	if err != nil {
		return nil, err
	}

	user, err := u.userService.UpdateUser(ctx, data, id)

	if err != nil {
		return nil, err
	}

	return user, nil

}
