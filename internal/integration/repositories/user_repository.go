package repositories

import (
	"context"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/integration/builder"
	"go-auth-api/internal/integration/models"
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
func (ur *userRepositoryImpl) Insert(ctx context.Context, data models.UsersModel) error {

	err := ur.db.WithContext(ctx).Model(&ur.model).Create(&data)

	if err != nil {

		return err.Error
	}

	return err.Error

}

func (ur *userRepositoryImpl) Update(ctx context.Context, data models.UsersModel) error {

	err := ur.db.WithContext(ctx).Model(&ur.model).Save(&data)

	if err != nil {

		return err.Error
	}

	return nil

}
func (ur *userRepositoryImpl) Select(ctx context.Context, queryParams dtos.QueryParams) (*[]models.UsersModel, error) {
	var users []models.UsersModel

	query := builder.BuildGormQuery(queryParams)

	pagination := builder.BuildGormPagination(queryParams)

	err := ur.db.WithContext(ctx).Model(&ur.model).Scopes(query...).Scopes(pagination).Find(&users)

	if err != nil {
		return nil, err.Error
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
