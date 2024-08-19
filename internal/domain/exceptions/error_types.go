package exceptions

type ErrorType interface {
	Error() string
	Description() string
	ToMap() *map[string]interface{}
	StatusCode() int
	Details() []ErrorDetails
}

type errorType struct {
	id          string
	statusCode  int
	code        string
	details     []ErrorDetails
	description string
	message     string
}

type ErrorDetails struct {
	Attribute string `json:"attribute"`
	Messages  string `json:"messages"`
}

func NewErrorDetail(attribute string, message string) ErrorDetails {

	return ErrorDetails{
		Attribute: attribute,
		Messages:  message,
	}
}

func (e *errorType) ToMap() *map[string]interface{} {
	result := map[string]interface{}{
		"id":          e.id,
		"code":        e.code,
		"description": e.description,
		"message":     e.message,
	}

	if e.details != nil {
		result["details"] = e.details
	}
	return &result
}

func (e *errorType) Error() string {
	return e.message
}

func (e *errorType) Description() string {
	return e.description
}

func (e *errorType) StatusCode() int {
	return e.statusCode
}

func (e *errorType) Details() []ErrorDetails { return e.details }

type HttpException struct {
	Id          string         `json:"id"`
	Description string         `json:"description"`
	Datetime    string         `json:"date_time"`
	Details     []ErrorDetails `json:"details,omitempty"`
	StatusCode  int            `json:"status_code"`
}
