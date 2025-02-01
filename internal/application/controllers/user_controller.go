package controllers

import (
	"context"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"go-auth-api/internal/application/parsers"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"go-auth-api/internal/domain/types"
	"net/http"
)

type userController struct {
	createUser adapters.CreateUserUseCase
	updateUser adapters.UpdateUserUseCaseAdapter
	findUser   adapters.FindUserUseCaseAdapter
}

func UserControllerConstructor(
	createUser adapters.CreateUserUseCase,
	updateUser adapters.UpdateUserUseCaseAdapter,
	findUser adapters.FindUserUseCaseAdapter,
) adapters.UsersController {
	return &userController{
		updateUser: updateUser,
		createUser: createUser,
		findUser:   findUser,
	}
}

func (c *userController) Post(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "request_id", uuid.New().String())
	var err error
	var payload *dtos.UsersDto

	err = render.DecodeJSON(r.Body, &payload)

	if err != nil {
		errResponse := parsers.HttpErrorParser(nil, ctx, err)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}
	result, errBusiness := c.createUser.Execute(ctx, payload)

	if errBusiness != nil {

		errResponse := parsers.HttpErrorParser(errBusiness, ctx, nil)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	render.Status(r, 201)
	render.JSON(rw, r, map[string]interface{}{"data": *result})
}

func (c *userController) Put(rw http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	ctx = context.WithValue(ctx, "request_id", uuid.New().String())
	var err error
	var payload *dtos.UpdateUserDto

	err = render.DecodeJSON(r.Body, &payload)

	if err != nil {
		errResponse := parsers.HttpErrorParser(nil, ctx, err)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	id := chi.URLParam(r, "id")

	parsedUuid, err := uuid.Parse(id)

	if err != nil {
		errResponse := parsers.HttpErrorParser(nil, ctx, err)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	result, errBusiness := c.updateUser.Execute(ctx, payload, parsedUuid)

	if errBusiness != nil {
		errResponse := parsers.HttpErrorParser(errBusiness, ctx, nil)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	render.Status(r, 200)
	render.JSON(rw, r, map[string]interface{}{"data": *result})
}

func (c *userController) List(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "request_id", uuid.New().String())

	var query dtos.QueryParams

	params, errQuery := query.BuildQuery(ctx, r.URL.Query())
	if errQuery != nil {
		errResponse := parsers.HttpErrorParser(errQuery, ctx, nil)
		render.Status(r, 500)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}
	result, err := c.findUser.ListUser(ctx, *params)
	if err != nil {
		errResponse := parsers.HttpErrorParser(err, ctx, nil)
		render.Status(r, 400)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	totalItems := len(*result)
	page := 1
	limit := 20
	if query.Page != nil {
		page = *query.Page
	}
	if query.Limit != nil {
		limit = *query.Limit
	}
	totalPages := (totalItems + limit - 1) / limit

	var data []interface{}
	for _, user := range *result {
		data = append(data, user)
	}
	response := types.Pagination{
		Data:        data,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}
	render.Status(r, 200)
	render.JSON(rw, r, response)

}

func (c *userController) FindById(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "request_id", uuid.New().String())

	id := chi.URLParam(r, "id")
	userId, err := uuid.Parse(id)
	if err != nil {
		errResponse := parsers.HttpErrorParser(nil, ctx, err)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	result, errBusiness := c.findUser.FindUser(ctx, userId)

	if errBusiness != nil {

		errResponse := parsers.HttpErrorParser(errBusiness, ctx, nil)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	render.Status(r, 201)
	render.JSON(rw, r, map[string]interface{}{"data": *result})
}
