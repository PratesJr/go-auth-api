package adapters

import (
	"context"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/integration/models"
)

type UserRepository interface {
	Insert(ctx context.Context, data models.UsersModel) *dtos.Exception
	Update(ctx context.Context, data models.UsersModel) *dtos.Exception
	Select(ctx context.Context, queryParams dtos.QueryParams) (*dtos.Exception, *[]models.UsersModel)
	Count(ctx context.Context, queryParams dtos.QueryParams) (*dtos.Exception, *int64)
}
