package services

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
	"go-auth-api/internal/utils"
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
	return u.userPersistence.Update(ctx, data, id.String())
}

func (u userServiceImpl) ListUser(ctx context.Context, params dtos.QueryParams) (*[]types.User, exceptions.ErrorType) {
	return u.userPersistence.Find(ctx, params)
}

func (u userServiceImpl) FindUserById(ctx context.Context, id *uuid.UUID) (*types.User, exceptions.ErrorType) {
	params := dtos.QueryParams{
		Id: utils.ToPointer(id.String()),
	}
	users, err := u.userPersistence.Find(ctx, params)

	if len(*users) == 0 {

		return nil, nil
	}

	var user *types.User

	for _, value := range *users {

		*user = types.User{
			Id:        value.Id,
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt,
		}

	}

	return user, err

}
