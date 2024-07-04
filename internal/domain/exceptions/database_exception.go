package exceptions

import (
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/enums"
)

func DatabaseException(message string, err error) dtos.Exception {
	return dtos.Exception{
		Message: message,
		Type:    enums.DatabaseError,
		Error:   err,
	}
}
