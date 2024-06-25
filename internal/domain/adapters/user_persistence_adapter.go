package adapters

import (
	"context"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
)

type UserPersistence interface {
	Create(ctx context.Context, data *dtos.UsersDto) (error, *types.User)
	Find(ctx context.Context, params *dtos.QueryParams) (error, *[]types.User)
	Update(ctx context.Context, data *dtos.UpdateUserDto, id string) (error, *types.User)
}
