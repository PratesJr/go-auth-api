package exceptions

import (
	"go-auth-api/internal/domain/enums"
	"go-auth-api/internal/domain/types"
)

func DuplicatedUserException(message string, err error) types.Exception {
	return types.Exception{
		Message: message,
		Type:    enums.DuplicatedError,
		Error:   err,
	}
}
