package types

type User struct {
	Id        *string `json:"id"`
	Name      *string `json:"name"`
	Email     *string `json:"email"`
	Birth     *string `json:"birth"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}
