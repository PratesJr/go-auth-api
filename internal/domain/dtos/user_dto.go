package dtos

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UsersDto struct {
	Name     string    `json:"name"`
	Email    string    `json:"email" validate:"email"`
	Birth    time.Time `json:"birth"`
	Password string    `json:"password"`
}

func NewUserDto(name string, email string, birth time.Time, password string) (*UsersDto, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	return &UsersDto{
		Name:     name,
		Email:    email,
		Birth:    birth,
		Password: string(bytes),
	}, nil
}
