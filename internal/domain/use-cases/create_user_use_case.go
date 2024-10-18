package use_cases

import (
	"context"
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
	//TODO implement me
	panic("implement me")
}
