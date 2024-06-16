package models

import (
	"github.com/google/uuid"
	"time"
)

type UsersModel struct {
	ID        *uuid.UUID `json:"id" gorm:"type: uuid; not null; primaryKey"`
	Name      *string    `json:"name" gorm:"type: varchar(100); not null"`
	Email     *string    `json:"email" validate:"email" gorm:"type: varchar(100); not null"`
	Birth     *time.Time `json:"birth" gorm:"type: date; not null"`
	Password  *string    `json:"password" gorm:"type: varchar; not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"type: timestamp; not null"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"type: timestamp; not null"`
}

func (UsersModel) TableName() string {
	return "auth_users.users"
}
