package adapters

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type UpdateUserUseCaseAdapter interface {
	Execute(ctx context.Context, data *dtos.UpdateUserDto, id uuid.UUID) (*types.User, exceptions.ErrorType)
}
