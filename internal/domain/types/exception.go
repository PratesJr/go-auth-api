package types

import "go-auth-api/internal/domain/enums"

type Exception struct {
	Message string
	Type    enums.ErrorType
	Error   error
}
