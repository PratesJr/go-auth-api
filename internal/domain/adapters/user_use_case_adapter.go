package adapters

import (
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
)

type UserUseCase interface {
	Create(user *dtos.UsersDto) (error, *types.User)
	Update(data *dtos.UpdateUserDto) (error, *types.User)
	Find() (error, *[]types.User)
}
