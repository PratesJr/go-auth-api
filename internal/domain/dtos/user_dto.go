package dtos

import "github.com/google/uuid"

type Users struct {
	ID *uuid.UUID `json:"id" gorm:"type: uuid; not null; primaryKey"`
}

func (Users) TableName() string {
	return "postgres.user"
}
