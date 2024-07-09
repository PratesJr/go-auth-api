package adapters

import (
	"context"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
)

type UserPersistence interface {
	Create(ctx context.Context, data *dtos.UsersDto) (*types.Exception, *types.User)
	Find(ctx context.Context, params *dtos.QueryParams) (*types.Exception, *[]types.User)
	Update(ctx context.Context, data *dtos.UpdateUserDto, id string) (*types.Exception, *types.User)
}
