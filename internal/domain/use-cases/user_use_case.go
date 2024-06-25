package use_cases

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
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

func (uc *userUseCaseImpl) Create(ctx context.Context, user *dtos.UsersDto) (error, *types.User) {
	return nil, nil
}

func (uc *userUseCaseImpl) Update(ctx context.Context, data *dtos.UpdateUserDto, id uuid.UUID) (error, *types.User) {
	return nil, nil
}

func (uc *userUseCaseImpl) Find(ctx context.Context, params dtos.QueryParams) (error, *[]types.User) {
	return nil, nil
}
