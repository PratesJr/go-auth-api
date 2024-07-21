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
}

var StatusCode = &httpStatus{
	BadRequest:          400,
	UnprocessableEntity: 422,
	Forbidden:           403,
	Conflict:            409,
	Unauthorized:        401,
	InternalServerError: 500,
	UnknownError:        520,
	BadGateway:          502,
}
