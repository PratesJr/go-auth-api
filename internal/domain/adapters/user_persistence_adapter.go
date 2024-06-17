package adapters

import (
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
)

type UserPersistence interface {
	Create(data *dtos.UsersDto) (error, *types.User)
	Find() (error, *[]types.User)
	Update(data *dtos.UsersDto) (error, *types.User)
}
