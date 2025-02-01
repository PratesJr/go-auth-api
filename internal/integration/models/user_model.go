package models

import (
	"github.com/google/uuid"
	"go-auth-api/internal/domain/dtos"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UsersModel struct {
	ID        uuid.UUID `gorm:"type:uuid; not null; primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100); not null" json:"name"`
	Email     string    `gorm:"type:varchar(100); not null" json:"email"`
	Birth     time.Time `gorm:"type:date; not null" json:"birth"`
	Password  string    `gorm:"type:varchar; not null" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp; not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; not null" json:"updated_at"`
}

func (UsersModel) TableName() string {
	return "auth_users.users"
}

func NewUserModel(body dtos.UsersDto) *UsersModel {
	parseDate, err := time.Parse(time.DateOnly, *body.Birth)

	if err != nil {
		return nil
	}

	bytes, errPassword := bcrypt.GenerateFromPassword([]byte(*body.Password), 14)

	if errPassword != nil {
		return nil
	}

	return &UsersModel{
		ID:        uuid.New(),
		Name:      *body.Name,
		Email:     *body.Email,
		Birth:     parseDate,
		Password:  string(bytes),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func UpdateUserData(dto *dtos.UpdateUserDto, user *UsersModel) *UsersModel {

	if dto.Birth != nil {
		parseDate, err := time.Parse(time.DateOnly, *dto.Birth)

		if err != nil {
			return nil
		}
		user.Birth = parseDate
	}

	if dto.Name != nil {
		user.Name = *dto.Name

	}

	if dto.Email != nil {
		user.Email = *dto.Email
	}

	user.UpdatedAt = time.Now().UTC()

	return user
}

func (user UsersModel) UpdateData(dto dtos.UpdateUserDto) *UsersModel {

	if dto.Birth != nil {
		parseDate, err := time.Parse("YYYY-MM-DD", *dto.Birth)

		if err != nil {
			return nil
		}
		user.Birth = parseDate
	}

	if dto.Name != nil {
		user.Name = *dto.Name

	}

	if dto.Email != nil {
		user.Email = *dto.Email
	}

	user.UpdatedAt = time.Now().UTC()

	return &user
}
