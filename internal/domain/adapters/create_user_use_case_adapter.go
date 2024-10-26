package adapters

import (
	"context"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type CreateUserUseCase interface {
	Execute(ctx context.Context, user *dtos.UsersDto) (*types.User, exceptions.ErrorType)
}
