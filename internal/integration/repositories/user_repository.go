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
	model *models.UsersModel
	db    *gorm.DB
}

func UserRepositoryConstructor(database *gorm.DB, model *models.UsersModel) adapters.UserRepository {
	return &userRepositoryImpl{
		model: model,
		db:    database,
	}
}
func (ur *userRepositoryImpl) Insert(ctx context.Context, data models.UsersModel) error {

	user := ur.db.WithContext(ctx).Model(&ur.model).Create(&data)

	return user.Error

}

func (ur *userRepositoryImpl) Update(ctx context.Context, data models.UsersModel) error {

	user := ur.db.WithContext(ctx).Model(&ur.model).Save(&data)

	return user.Error

}
func (ur *userRepositoryImpl) Select(ctx context.Context, queryParams dtos.QueryParams) (error, *[]models.UsersModel) {
	var users []models.UsersModel

	query := builder.BuildGormQuery(queryParams)

	pagination := builder.BuildGormPagination(queryParams)

	err := ur.db.WithContext(ctx).Model(&ur.model).Scopes(query...).Scopes(pagination).Find(&users)

	if err != nil {
		return err.Error, nil
	}
	return nil, &users

}
func (ur *userRepositoryImpl) Count(ctx context.Context, queryParams dtos.QueryParams) (error, *int64) {
	var count *int64

	query := builder.BuildGormQuery(queryParams)

	ur.db.WithContext(ctx).Model(&ur.model).Scopes(query...).Count(count)

	return nil, count
}
