package adapters

import (
	"context"
	"github.com/google/uuid"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
)

type FindUserUseCaseAdapter interface {
	ListUser(ctx context.Context, params dtos.QueryParams) (*[]types.User, exceptions.ErrorType)
	FindUser(ctx context.Context, id uuid.UUID) (*types.User, exceptions.ErrorType)
}
