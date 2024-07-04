package enums

type ErrorType struct {
	Description string
	StatusCode  HttpStatus
}

var ForbiddenError = ErrorType{
	Description: "Forbidden.",
	StatusCode:  Forbidden,
}
var UnauthorizedError = ErrorType{
	Description: "No authorization provided.",
	StatusCode:  Unauthorized,
}

var UnknownError = ErrorType{
	Description: "There is something wrong on api.",
	StatusCode:  ServerInternalError,
}

var DatabaseError = ErrorType{
	Description: "Error on PostgreSQL database.",
	StatusCode:  ServerInternalError,
}

var RequestError = ErrorType{
	Description: "Request Error.",
	StatusCode:  BadRequest,
}

var DuplicatedError = ErrorType{
	Description: "Unable to create, duplicated data.",
	StatusCode:  Conflict,
}
var UnprocessableEntityError = ErrorType{
	Description: "Invalid request body.",
	StatusCode:  UnprocessableEntity,
}
