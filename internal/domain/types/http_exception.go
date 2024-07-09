package types

type HttpException struct {
	Id          string    `json:"id"`
	Description string    `json:"description"`
	Datetime    string    `json:"date_time"`
	Messages    *[]string `json:"messages,omitempty"`
}
