package adapters

import "go-auth-api/internal/integration/models"

type UserRepository interface {
	Insert(data *models.UsersModel) error
	Update(data *models.UsersModel) error
	Select() (error, []models.UsersModel)
}
