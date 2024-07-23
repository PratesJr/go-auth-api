package exceptions

import (
	"context"
	"encoding/json"
	"go-auth-api/internal/application/config"
	"go-auth-api/internal/domain/enums"
)

func BadRequestException(ctx context.Context, errMessage string, details ...ErrorDetails) ErrorType {
	id := config.GetRequestId(ctx)

	var mapDetails []map[string]interface{}
	if details != nil {

		a, _ := json.Marshal(details)
		_ = json.Unmarshal(a, &mapDetails)
	}

	ex := errorType{
		id:          id,
		statusCode:  enums.StatusCode.BadRequest,
		code:        "400_BAD_REQUEST",
		details:     mapDetails,
		description: "Bad Request.",
		message:     errMessage,
	}
	return &ex
}

func UnauthorizedException(ctx context.Context, errMessage string) ErrorType {
	id := config.GetRequestId(ctx)

	ex := errorType{
		id:          id,
		statusCode:  enums.StatusCode.Unauthorized,
		code:        "401_UNAUTHORIZED",
		details:     nil,
		description: "Unauthorized.",
		message:     errMessage,
	}
	return &ex
}
