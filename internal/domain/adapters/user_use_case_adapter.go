package adapters

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
)

type UserUseCase interface {
	Create(ctx context.Context, user *dtos.UsersDto) (error, *types.User)
	Update(ctx context.Context, data *dtos.UpdateUserDto, id uuid.UUID) (error, *types.User)
	Find(ctx context.Context, params dtos.QueryParams) (error, *[]types.User)
}
