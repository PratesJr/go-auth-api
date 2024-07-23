package adapters

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type UserUseCase interface {
	Create(ctx context.Context, user *dtos.UsersDto) (*types.User, *exceptions.ErrorType)
	Update(ctx context.Context, data *dtos.UpdateUserDto, id uuid.UUID) (*types.User, *exceptions.ErrorType)
	Find(ctx context.Context, params dtos.QueryParams) (*[]types.User, *exceptions.ErrorType)
}
