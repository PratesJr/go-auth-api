package use_cases

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type userUseCaseImpl struct {
	userPersistence adapters.UserPersistence
}

func UserUseCaseConstructor(userPersistence adapters.UserPersistence) adapters.UserUseCase {
	return &userUseCaseImpl{
		userPersistence: userPersistence,
	}
}

func (uc *userUseCaseImpl) Create(ctx context.Context, user *dtos.UsersDto) (*types.User, *exceptions.ErrorType) {

	createdUser, err := uc.userPersistence.Create(ctx, user)

	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (uc *userUseCaseImpl) Update(ctx context.Context, data *dtos.UpdateUserDto, id uuid.UUID) (*types.User, *exceptions.ErrorType) {
	return nil, nil
}

func (uc *userUseCaseImpl) Find(ctx context.Context, params dtos.QueryParams) (*[]types.User, *exceptions.ErrorType) {
	return nil, nil
}
