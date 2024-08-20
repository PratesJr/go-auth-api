package exceptions

import (
	"context"
	"go-auth-api/internal/application/config"
	"go-auth-api/internal/domain/enums"
)

func BadRequestException(ctx context.Context, errMessage string, details ...ErrorDetails) ErrorType {
	id := config.GetRequestId(ctx)

	ex := errorType{
		id:          id,
		statusCode:  enums.StatusCode.BadRequest,
		code:        "400_BAD_REQUEST",
		details:     details,
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

func UnprocessableEntity(ctx context.Context, errMessage string, details []ErrorDetails) ErrorType {
	id := config.GetRequestId(ctx)

	ex := errorType{
		id:          id,
		statusCode:  enums.StatusCode.UnprocessableEntity,
		code:        "422_UNPROCESSABLE_ENTITY",
		details:     details,
		description: "Unprocessable entity.",
		message:     errMessage,
	}
	return &ex
}

func DatabaseException(ctx context.Context, errMessage string) ErrorType {
	id := config.GetRequestId(ctx)

	ex := errorType{
		id:          id,
		statusCode:  enums.StatusCode.InternalServerError,
		code:        "500_INTERNAL_ERROR",
		details:     nil,
		description: "Something went wrong.",
		message:     errMessage,
	}
	return &ex
}
