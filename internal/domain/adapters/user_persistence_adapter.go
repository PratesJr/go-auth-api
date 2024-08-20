package adapters

import (
	"context"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type UserPersistence interface {
	Create(ctx context.Context, data *dtos.UsersDto) (*types.User, exceptions.ErrorType)
	Find(ctx context.Context, params dtos.QueryParams) (*[]types.User, exceptions.ErrorType)
	Update(ctx context.Context, data dtos.UpdateUserDto, id string) (*types.User, exceptions.ErrorType)
}
