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
			return db.Where("email = ?", params.Email)
		})
	}

	if params.Id != nil {
		query = append(query, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", params.Id)
		})
	}

	return query
}

func BuildGormPagination(params dtos.QueryParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := 1
		limit := 10
		orderBy := "id ASC"

		if params.Page != nil {
			page = *params.Page
		}
		if params.Limit != nil {
			limit = *params.Limit
		}
		if params.OrderBy != nil {
			orderBy = *params.OrderBy
		}

		return db.
			Offset((page - 1) * limit).
			Limit(limit).
			Order(orderBy)
	}
}
