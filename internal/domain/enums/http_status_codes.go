package enums

type httpStatus struct {
	BadRequest          int
	UnprocessableEntity int
	Forbidden           int
	Conflict            int
	Unauthorized        int
	InternalServerError int
	UnknownError        int
	BadGateway          int
	NotFound            int
}

var StatusCode = &httpStatus{
	BadRequest:          400,
	Unauthorized:        401,
	Forbidden:           403,
	NotFound:            404,
	Conflict:            409,
	UnprocessableEntity: 422,
	InternalServerError: 500,
	UnknownError:        520,
	BadGateway:          502,
}
