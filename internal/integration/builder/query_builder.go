package builder

import (
	"go-auth-api/internal/domain/dtos"
	"gorm.io/gorm"
)

func BuildGormQuery(params dtos.QueryParams) []func(db *gorm.DB) *gorm.DB {
	var query []func(db *gorm.DB) *gorm.DB

	if params.Name != nil {
		query = append(query, func(db *gorm.DB) *gorm.DB {
			return db.Where("name LIKE ?", *params.Name+"%")
		})

	}

	if params.Email != nil {
		query = append(query, func(db *gorm.DB) *gorm.DB {
			return db.Where("email = ?", *params.Email)
		})
	}

	if params.Id != nil {
		query = append(query, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", *params.Id)
		})
	}

	return query
}

func BuildGormPagination(params dtos.QueryParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.
			Offset((*params.Page - 1) * *params.Limit).
			Limit(*params.Limit).
			Order(*params.OrderBy)
	}
}
