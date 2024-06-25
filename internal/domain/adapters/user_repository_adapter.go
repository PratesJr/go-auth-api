package adapters

import (
	"go-auth-api/internal/integration/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Insert(data *models.UsersModel) error
	Update(data *models.UsersModel) error
	Select(query []func(db *gorm.DB) *gorm.DB) (error, []models.UsersModel)
	Count(query []func(db *gorm.DB) *gorm.DB) (error, *int64)
}
