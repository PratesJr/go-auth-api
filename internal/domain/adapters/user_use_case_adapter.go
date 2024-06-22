package adapters

import (
	"github.com/google/uuid"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
)

type UserUseCase interface {
	Create(user *dtos.UsersDto) (error, *types.User)
	Update(data *dtos.UpdateUserDto, id uuid.UUID) (error, *types.User)
	Find() (error, *[]types.User)
}
