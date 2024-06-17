package persistences

import (
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
	"go-auth-api/internal/integration/models"
)

type UserPersistenceImpl struct {
	repo adapters.UserRepository
}

func (up UserPersistenceImpl) Create(data *dtos.UsersDto) (error, *types.User) {
	dataToPersist := models.NewUserModel(data)

	err := up.repo.Insert(dataToPersist)

	if err != nil {
		return err, nil
	}

	return nil, &types.User{
		Id:        dataToPersist.ID.String(),
		Name:      dataToPersist.Name,
		Email:     dataToPersist.Email,
		CreatedAt: dataToPersist.CreatedAt.String(),
	}
}

func (up UserPersistenceImpl) Find() (error, *[]types.User) {
	return nil, nil
}

func (up UserPersistenceImpl) Update() (error, *types.User) {
	err, user := up.Find()

	if err != nil {
		return err, nil
	}
	return nil, nil
}
