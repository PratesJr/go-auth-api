package dtos

type UsersDto struct {
	Name     *string `json:"name" validate:"required"`
	Email    *string `json:"email" validate:"required,email"`
	Birth    *string `json:"birth" validate:"required"`
	Password *string `json:"password" validate:"required"`
}
