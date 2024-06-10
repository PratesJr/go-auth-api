package repositories

import (
	"go-auth-api/internal/integration/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (ur *UserRepositoryImpl) Insert(data models.User) error {

}

func (ur *UserRepositoryImpl) Update(data models.User) error {

}

func (ur *UserRepositoryImpl) Select(data models.User) ([]models.User, error) {

}
