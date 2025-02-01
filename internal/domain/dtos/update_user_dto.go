package dtos

type UpdateUserDto struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty" validate:"omitempty,email"`
	Birth    *string `json:"birth,omitempty"`
	Password *string `json:"password,omitempty"`
}
