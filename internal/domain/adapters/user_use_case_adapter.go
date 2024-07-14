package adapters

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
)

type UserUseCase interface {
	Create(ctx context.Context, user *dtos.UsersDto) (*types.Exception, *types.User)
	Update(ctx context.Context, data *dtos.UpdateUserDto, id uuid.UUID) (*types.Exception, *types.User)
	Find(ctx context.Context, params dtos.QueryParams) (*types.Exception, *[]types.User)
}
