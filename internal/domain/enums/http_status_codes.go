package enums

type HttpStatus int

const (
	Forbidden           HttpStatus = 403
	Unauthorized        HttpStatus = 401
	Conflict            HttpStatus = 409
	BadRequest          HttpStatus = 400
	ServerInternalError HttpStatus = 500
	UnprocessableEntity HttpStatus = 422
)
