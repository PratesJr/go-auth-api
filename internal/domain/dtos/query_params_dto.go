package dtos

import (
	"go-auth-api/internal/utils"
	"net/url"
	"strconv"
	"strings"
)

type QueryParams struct {
	Id      *string `json:"id,omitempty" validate:"uuid"`
	Email   *string `json:"email,omitempty" validate:"email"`
	Name    *string `json:"name,omitempty"`
	Page    *int    `json:"page,omitempty"`
	Limit   *int    `json:"limit,omitempty"`
	OrderBy *string `json:"order,omitempty"`
}

func (q *QueryParams) BuildQuery(urlQuery url.Values) *QueryParams {

	q = &QueryParams{
		Id:      utils.NilIfIsEmpty(urlQuery.Get("id")),
		Email:   utils.NilIfIsEmpty(urlQuery.Get("email")),
		Name:    utils.NilIfIsEmpty(urlQuery.Get("name")),
		Page:    nil,
		Limit:   nil,
		OrderBy: utils.NilIfIsEmpty(urlQuery.Get("order")),
	}

	limit, _ := strconv.Atoi(urlQuery.Get("limit"))

	if limit == 0 {
		limit = 20
	}

	q.Limit = utils.ToPointer(limit)

	page, _ := strconv.Atoi(urlQuery.Get("page"))

	if page == 0 {
		page = 1
	}

	q.Page = utils.ToPointer(page)

	if urlQuery.Get("order") != "" {

		if strings.Contains(urlQuery.Get("order"), "-") {
			splitStr := strings.Split(urlQuery.Get("order"), "-")

			q.OrderBy = utils.ToPointer(splitStr[0] + " " + "ASC")
		}

		if !strings.Contains(urlQuery.Get("order"), "-") {
			q.OrderBy = utils.ToPointer(urlQuery.Get("order") + " " + "DESC")
		}

		return q
	}

	q.OrderBy = utils.ToPointer("created_at DESC")

	return q
}
