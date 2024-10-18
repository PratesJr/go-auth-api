package adapters

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type UserService interface {
	NewUser(ctx context.Context, user *dtos.UsersDto) (*types.User, exceptions.ErrorType)
	UpdateUser(ctx context.Context, data *dtos.UpdateUserDto, id uuid.UUID) (*types.User, exceptions.ErrorType)
	ListUser(ctx context.Context, params dtos.QueryParams) (*[]types.User, exceptions.ErrorType)
	FindUserById(ctx context.Context, id *uuid.UUID) (*types.User, exceptions.ErrorType)
}
