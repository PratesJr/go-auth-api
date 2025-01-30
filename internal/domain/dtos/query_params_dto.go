package dtos

import (
	"context"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/utils"
	"net/url"
	"strconv"
	"strings"
)

type QueryParams struct {
	Id      *string `json:"id,omitempty" validate:"omitempty,uuid"`
	Email   *string `json:"email,omitempty" validate:"omitempty,email"`
	Name    *string `json:"name,omitempty"`
	Page    *int    `json:"page,omitempty"`
	Limit   *int    `json:"limit,omitempty"`
	OrderBy *string `json:"order,omitempty"`
}

func (q *QueryParams) BuildQuery(ctx context.Context, urlQuery url.Values) (*QueryParams, exceptions.ErrorType) {
	id := urlQuery.Get("id")
	if id != "" && !utils.IsValidUUID(id) {
		return nil, exceptions.BadRequestException(ctx, "invalid UUID format for 'id'")
	}

	query := &QueryParams{
		Id:      utils.NilIfIsEmpty(id),
		Email:   utils.NilIfIsEmpty(urlQuery.Get("email")),
		Name:    utils.NilIfIsEmpty(urlQuery.Get("name")),
		Page:    nil,
		Limit:   nil,
		OrderBy: nil,
	}

	limitStr := urlQuery.Get("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			return nil, exceptions.BadRequestException(ctx, "invalid 'limit' value")
		}
		query.Limit = utils.ToPointer(limit)
	} else {
		query.Limit = utils.ToPointer(20)
	}

	pageStr := urlQuery.Get("page")
	if pageStr != "" {
		page, err := strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			return nil, exceptions.BadRequestException(ctx, "invalid 'page' value")
		}
		query.Page = utils.ToPointer(page)
	} else {
		query.Page = utils.ToPointer(1)
	}

	order := urlQuery.Get("order")
	if order != "" {
		if strings.HasPrefix(order, "-") {
			query.OrderBy = utils.ToPointer(strings.TrimPrefix(order, "-") + " DESC")
		} else {
			query.OrderBy = utils.ToPointer(order + " ASC")
		}
	} else {
		query.OrderBy = utils.ToPointer("created_at DESC")
	}

	return query, nil
}
