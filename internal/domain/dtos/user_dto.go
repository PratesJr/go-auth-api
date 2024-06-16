package dtos

import "github.com/google/uuid"

type Users struct {
	ID   *uuid.UUID `json:"id" gorm:"type: uuid; not null; primaryKey"`
	Name *string    `json:"name" gorm:"type: varchar(100); not null"`
}

func (Users) TableName() string {
	return "auth_users.users"
}
