package parsers

import (
	"context"
	"go-auth-api/internal/domain/enums"
	"go-auth-api/internal/domain/exceptions"
	"time"
)

func HttpErrorParser(customErr exceptions.ErrorType, ctx context.Context, err error) exceptions.HttpException {
	var errResponse exceptions.HttpException

	errResponse.Datetime = time.Now().Format(time.DateTime)
	errResponse.Id = ctx.Value("request_id").(string)

	if customErr == nil && err != nil {
		errResponse.StatusCode = enums.StatusCode.InternalServerError
		errResponse.Description = "Unexpected error."

	}
	if customErr != nil {
		errResponse.StatusCode = customErr.StatusCode()
		errResponse.Description = customErr.Description()
	}

	if customErr.Details() != nil {
		errResponse.Details = customErr.Details()
	}
	return errResponse
}
