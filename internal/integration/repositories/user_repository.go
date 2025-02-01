package repositories

import (
	"context"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/integration/builder"
	"go-auth-api/internal/integration/models"
	"go-auth-api/internal/utils"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	model models.UsersModel
	db    gorm.DB
}

func UserRepositoryConstructor(database gorm.DB, model models.UsersModel) adapters.UserRepository {
	return &userRepositoryImpl{
		model: model,
		db:    database,
	}
}
func (ur *userRepositoryImpl) Insert(ctx context.Context, data *models.UsersModel) error {

	err := ur.db.WithContext(ctx).Model(&ur.model).Create(data)

	if err != nil {

		return err.Error
	}

	return err.Error

}

func (ur *userRepositoryImpl) Update(ctx context.Context, data *models.UsersModel) error {

	queryParams := dtos.QueryParams{Id: utils.ToPointer(data.ID.String())}

	query := builder.BuildGormQuery(queryParams)

	err := ur.db.WithContext(ctx).Model(&ur.model).Scopes(query...).Save(data)

	if err != nil {

		return err.Error
	}

	return nil

}
func (ur *userRepositoryImpl) Select(ctx context.Context, queryParams dtos.QueryParams) (*[]models.UsersModel, error) {
	var users []models.UsersModel

	query := builder.BuildGormQuery(queryParams)
	pagination := builder.BuildGormPagination(queryParams)

	if err := ur.db.WithContext(ctx).
		Model(&ur.model).
		Scopes(query...).
		Scopes(pagination).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (ur *userRepositoryImpl) Count(ctx context.Context, queryParams dtos.QueryParams) (*int64, error) {
	var count *int64

	query := builder.BuildGormQuery(queryParams)

	err := ur.db.WithContext(ctx).Model(&ur.model).Scopes(query...).Count(count)

	if err.Error != nil {
		return nil, err.Error
	}

	return count, nil
}
