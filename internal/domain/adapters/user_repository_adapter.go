package adapters

import "go-auth-api/internal/integration/models"

type UserRepository interface {
	Insert(u models.User) error
	Update(u models.User) error
	Select() ([]models.User, error)
}
