package adapters

import (
	"context"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/integration/models"
)

type UserRepository interface {
	Insert(ctx context.Context, data models.UsersModel) error
	Update(ctx context.Context, data models.UsersModel) error
	Select(ctx context.Context, queryParams dtos.QueryParams) (*[]models.UsersModel, error)
	Count(ctx context.Context, queryParams dtos.QueryParams) (*int64, error)
}
