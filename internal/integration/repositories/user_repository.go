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
func (ur userRepositoryImpl) Insert(data *models.UsersModel) error {

	user := ur.db.Table(ur.model.TableName()).Create(&data)

	return user.Error

}

func (ur userRepositoryImpl) Update(data *models.UsersModel) error {

	user := ur.db.Table(ur.model.TableName()).Save(&data)

	return user.Error

}

func (ur userRepositoryImpl) Select() (error, []models.UsersModel) {
	return nil, nil

}
