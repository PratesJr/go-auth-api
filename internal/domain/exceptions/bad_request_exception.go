package exceptions

import (
	"go-auth-api/internal/domain/enums"
	"go-auth-api/internal/domain/types"
)

func BadRequestException(message string, err error) types.Exception {
	return types.Exception{
		Message: message,
		Type:    enums.RequestError,
		Error:   err,
	}
}
