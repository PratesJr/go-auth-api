package exceptions

import "go-auth-api/internal/domain/enums"

type ErrorType struct {
	Description string
	Error       string
	StatusCode  int
}

func BadRequestException(description string, err string) ErrorType {
	return ErrorType{
		Description: description,
		StatusCode:  enums.StatusCode.BadRequest,
		Error:       err,
	}
}

func ConflictException(description string, err string) ErrorType {
	return ErrorType{
		Description: description,
		StatusCode:  enums.StatusCode.Conflict,
		Error:       err,
	}
}

func ForbiddenException(description string, err string) ErrorType {
	return ErrorType{
		Description: description,
		StatusCode:  enums.StatusCode.Forbidden,
		Error:       err,
	}
}

func UnauthorizedException(description string, err string) ErrorType {
	return ErrorType{
		Description: description,
		StatusCode:  enums.StatusCode.Unauthorized,
		Error:       err,
	}
}
func UnprocessableEntityException(description string, err string) ErrorType {
	return ErrorType{
		Description: description,
		StatusCode:  enums.StatusCode.UnprocessableEntity,
		Error:       err,
	}
}
func InternalServerErrorException(description string, err string) ErrorType {
	return ErrorType{
		Description: description,
		StatusCode:  enums.StatusCode.InternalServerError,
		Error:       err,
	}
}

func BadGatewayException(description string, err string) ErrorType {
	return ErrorType{
		Description: description,
		StatusCode:  enums.StatusCode.BadGateway,
		Error:       err,
	}
}

func UnknownErrorException(description string, err string) ErrorType {
	return ErrorType{
		Description: description,
		StatusCode:  enums.StatusCode.UnknownError,
		Error:       err,
	}
}

type HttpException struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Datetime    string    `json:"date_time"`
	Messages    *[]string `json:"messages,omitempty"`
	StatusCode  int
}
