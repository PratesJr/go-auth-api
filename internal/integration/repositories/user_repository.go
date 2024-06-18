package repositories

import (
	"go-auth-api/internal/integration/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Model models.UsersModel
	Db    *gorm.DB
}

func (ur *UserRepositoryImpl) Insert(data *models.UsersModel) error {

	user := ur.Db.Table(ur.Model.TableName()).Create(&data)

	return user.Error

}

func (ur *UserRepositoryImpl) Update(data *models.UsersModel) error {

	user := ur.Db.Table(ur.Model.TableName()).Save(&data)

	return user.Error

}

func (ur *UserRepositoryImpl) Select() (error, []models.UsersModel) {
	return nil, nil

}
