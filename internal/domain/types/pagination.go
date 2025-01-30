package types

type Pagination struct {
	Data        []interface{} `json:"data"`
	CurrentPage int           `json:"current_page"`
	TotalPages  int           `json:"total_pages"`
	TotalItems  int           `json:"total_items"`
}
