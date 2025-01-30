package persistences

import (
	"context"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
	"go-auth-api/internal/integration/models"
	"go-auth-api/internal/utils"
)

type userPersistenceImpl struct {
	repo adapters.UserRepository
}

func UserPersistenceConstructor(repository adapters.UserRepository) adapters.UserPersistence {
	return &userPersistenceImpl{
		repo: repository,
	}
}

func (up *userPersistenceImpl) Create(ctx context.Context, data *dtos.UsersDto) (*types.User, exceptions.ErrorType) {
	dataToPersist := models.NewUserModel(*data)

	err := up.repo.Insert(ctx, dataToPersist)

	if err != nil {
		return nil, exceptions.DatabaseException(ctx, err.Error())
	}

	return &types.User{
		Id:        utils.ToPointer(dataToPersist.ID.String()),
		Name:      &dataToPersist.Name,
		Email:     &dataToPersist.Email,
		CreatedAt: utils.ToPointer(dataToPersist.CreatedAt.String()),
	}, nil
}

func (up *userPersistenceImpl) Find(ctx context.Context, params dtos.QueryParams) (*[]types.User, exceptions.ErrorType) {

	arrayUser, err := up.repo.Select(ctx, params)

	if err != nil {
		return nil, exceptions.DatabaseException(ctx, err.Error())
	}

	var result = make([]types.User, len(*arrayUser))

	for i, value := range *arrayUser {

		result[i] = types.User{
			Id:        utils.ToPointer(value.ID.String()),
			Name:      &value.Name,
			Email:     &value.Email,
			Birth:     utils.ToPointer(value.Birth.String()),
			CreatedAt: utils.ToPointer(value.CreatedAt.String()),
			UpdatedAt: utils.ToPointer(value.UpdatedAt.String()),
		}

	}

	return &result, nil
}

func (up *userPersistenceImpl) Update(ctx context.Context, data *dtos.UpdateUserDto, id string) (*types.User, exceptions.ErrorType) {

	querySelect := dtos.QueryParams{
		Id:    utils.ToPointer(id),
		Limit: utils.ToPointer(1),
	}

	userArr, err := up.repo.Select(ctx, querySelect)

	if err != nil {
		return nil, exceptions.DatabaseException(ctx, err.Error())
	}

	user := *userArr

	userToUpdate := user[0].UpdateData(*data)

	errUpdate := up.repo.Update(ctx, userToUpdate)

	if errUpdate != nil {
		return nil, exceptions.DatabaseException(ctx, err.Error())
	}

	return &types.User{
		Id:        utils.ToPointer(userToUpdate.ID.String()),
		Name:      &userToUpdate.Name,
		Email:     &userToUpdate.Email,
		CreatedAt: utils.ToPointer(userToUpdate.CreatedAt.String()),
	}, nil
}
