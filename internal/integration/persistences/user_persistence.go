package persistences

import (
	"context"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
	"go-auth-api/internal/integration/models"
	"go-auth-api/internal/utils"
)

type userPersistenceImpl struct {
	repo adapters.UserRepository
}

func UserPersistenceConstructor(repository *adapters.UserRepository) adapters.UserPersistence {
	return &userPersistenceImpl{
		repo: *repository,
	}
}

func (up *userPersistenceImpl) Create(ctx context.Context, data *dtos.UsersDto) (error, *types.User) {
	dataToPersist := models.NewUserModel(data)

	err := up.repo.Insert(ctx, *dataToPersist)

	if err != nil {
		return err, nil
	}

	return nil, &types.User{
		Id:        utils.ToPointer(dataToPersist.ID.String()),
		Name:      &dataToPersist.Name,
		Email:     &dataToPersist.Email,
		CreatedAt: utils.ToPointer(dataToPersist.CreatedAt.String()),
	}
}

func (up *userPersistenceImpl) Find(ctx context.Context, params *dtos.QueryParams) (error, *[]types.User) {

	err, arrayUser := up.repo.Select(ctx, *params)
	if arrayUser != nil {
		return err, nil
	}

	var result = make([]types.User, len(*arrayUser))

	for i, value := range *arrayUser {

		result[i] = types.User{
			Id:        utils.ToPointer(value.ID.String()),
			Name:      &value.Name,
			Email:     &value.Email,
			CreatedAt: utils.ToPointer(value.CreatedAt.String()),
		}

	}

	return nil, &result
}

func (up *userPersistenceImpl) Update(ctx context.Context, data *dtos.UpdateUserDto, id string) (error, *types.User) {

	querySelect := dtos.QueryParams{
		Id:    &id,
		Limit: utils.ToPointer(1),
	}

	err, userArr := up.repo.Select(ctx, querySelect)

	if err != nil {
		return err, nil
	}
	user := *userArr
	userToUpdate := user[0].UpdateData(data)

	errUpdate := up.repo.Update(ctx, *userToUpdate)

	if errUpdate != nil {
		return errUpdate, nil
	}

	return nil, &types.User{
		Id:        utils.ToPointer(userToUpdate.ID.String()),
		Name:      &userToUpdate.Name,
		Email:     &userToUpdate.Email,
		CreatedAt: utils.ToPointer(userToUpdate.CreatedAt.String()),
	}
}
