package repositories

import (
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/integration/models"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	model *models.UsersModel
	db    *gorm.DB
}

func UserRepositoryConstructor(database *gorm.DB, model *models.UsersModel) adapters.UserRepository {
	return &userRepositoryImpl{
		model: model,
		db:    database,
	}
}
func (ur *userRepositoryImpl) Insert(data *models.UsersModel) error {

	user := ur.db.Table(ur.model.TableName()).Create(&data)

	return user.Error

}

func (ur *userRepositoryImpl) Update(data *models.UsersModel) error {

	user := ur.db.Table(ur.model.TableName()).Save(&data)

	return user.Error

}

func (ur *userRepositoryImpl) Select(query []func(db *gorm.DB) *gorm.DB) (error, []models.UsersModel) {
	var users []models.UsersModel

	err := ur.db.Table(ur.model.TableName()).Scopes(query...).Find(&users)

	if err != nil {
		return err.Error, nil
	}
	return nil, users

}

func (ur *userRepositoryImpl) Count(query []func(db *gorm.DB) *gorm.DB) (error, *int64) {
	var count *int64

	ur.db.Table(ur.model.TableName()).Scopes(query...).Count(count)

	return nil, count
}
