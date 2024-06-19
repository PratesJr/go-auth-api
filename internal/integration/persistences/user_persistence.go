package persistences

import (
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
	"go-auth-api/internal/integration/models"
)

type userPersistenceImpl struct {
	repo adapters.UserRepository
}

func UserPersistenceConstructor(repository adapters.UserRepository) adapters.UserPersistence {
	return &userPersistenceImpl{
		repo: repository,
	}
}

func (up userPersistenceImpl) Create(data *dtos.UsersDto) (error, *types.User) {
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

func (up userPersistenceImpl) Find() (error, *[]types.User) {
	err, arrayUser := up.repo.Select()
	if arrayUser != nil {
		return err, nil
	}

	var result = make([]types.User, len(arrayUser))

	for i, value := range arrayUser {

		result[i] = types.User{
			Id:        value.ID.String(),
			Name:      value.Name,
			Email:     value.Email,
			CreatedAt: value.CreatedAt.String(),
		}

	}

	return nil, &result
}

func (up userPersistenceImpl) Update(data *dtos.UpdateUserDto, id string) (error, *types.User) {
	err, user := up.repo.Select()

	if err != nil {
		return err, nil
	}

	userToUpdate := user[0].UpdateData(data)

	errUpdate := up.repo.Update(userToUpdate)

	if errUpdate != nil {
		return errUpdate, nil
	}
	return nil, &types.User{
		Id:        userToUpdate.ID.String(),
		Name:      userToUpdate.Name,
		Email:     userToUpdate.Email,
		CreatedAt: userToUpdate.CreatedAt.String(),
	}
}
