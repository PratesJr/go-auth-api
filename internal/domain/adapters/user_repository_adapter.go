package adapters

import (
	"context"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/types"
	"go-auth-api/internal/integration/models"
)

type UserRepository interface {
	Insert(ctx context.Context, data models.UsersModel) *types.Exception
	Update(ctx context.Context, data models.UsersModel) *types.Exception
	Select(ctx context.Context, queryParams dtos.QueryParams) (*types.Exception, *[]models.UsersModel)
	Count(ctx context.Context, queryParams dtos.QueryParams) (*types.Exception, *int64)
}
