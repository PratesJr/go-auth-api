package services

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type userServiceImpl struct {
	userPersistence adapters.UserPersistence
}

func UserServiceConstructor(userPersistence adapters.UserPersistence) adapters.UserService {
	return &userServiceImpl{
		userPersistence: userPersistence,
	}
}

func (u userServiceImpl) NewUser(ctx context.Context, user *dtos.UsersDto) (*types.User, exceptions.ErrorType) {
	return u.userPersistence.Create(ctx, user)
}

func (u userServiceImpl) UpdateUser(ctx context.Context, data *dtos.UpdateUserDto, id uuid.UUID) (*types.User, exceptions.ErrorType) {
	//TODO implement me
	panic("implement me")
}

func (u userServiceImpl) ListUser(ctx context.Context, params dtos.QueryParams) (*[]types.User, exceptions.ErrorType) {
	//TODO implement me
	panic("implement me")
}

func (u userServiceImpl) FindUserById(ctx context.Context, id *uuid.UUID) (*types.User, exceptions.ErrorType) {
	//TODO implement me
	panic("implement me")
}
