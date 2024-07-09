package parsers

import (
	"context"
	"go-auth-api/internal/domain/enums"
	"go-auth-api/internal/domain/types"
	"time"
)

func HttpErrorParser(customErr *types.Exception, ctx context.Context, err *error) (*types.HttpException, *enums.HttpStatus) {
	var errResponse types.HttpException
	var statusCode enums.HttpStatus

	errResponse.Datetime = time.Now().String()
	errResponse.Id = ctx.Value("request_id").(string)

	if customErr == nil && err != nil {
		statusCode = enums.HttpStatus(500)
		errResponse.Description = "Unexpected error."

	}
	if customErr != nil {
		statusCode = customErr.Type.StatusCode
		errResponse.Description = customErr.Message
	}
	return &errResponse, &statusCode
}
