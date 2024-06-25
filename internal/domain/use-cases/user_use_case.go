package use_cases

import (
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

func (uc *userUseCaseImpl) Create(user *dtos.UsersDto) (error, *types.User) {
	return nil, nil
}

func (uc *userUseCaseImpl) Update(data *dtos.UpdateUserDto, id uuid.UUID) (error, *types.User) {
	return nil, nil
}

func (uc *userUseCaseImpl) Find(params dtos.QueryParams) (error, *[]types.User) {
	return nil, nil
}
