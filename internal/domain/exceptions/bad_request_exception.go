package exceptions

import (
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/enums"
)

func BadRequestException(message string, err error) dtos.Exception {
	return dtos.Exception{
		Message: message,
		Type:    enums.RequestError,
		Error:   err,
	}
}
