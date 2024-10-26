package controllers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"go-auth-api/internal/application/parsers"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
	"go-auth-api/internal/domain/exceptions"
	"net/http"
)

type userController struct {
	createUser adapters.CreateUserUseCase
	updateUser adapters.UpdateUserUseCaseAdapter
}

func UserControllerConstructor(
	createUser adapters.CreateUserUseCase,
	updateUser adapters.UpdateUserUseCaseAdapter,
) adapters.UsersController {
	return &userController{
		updateUser: updateUser,
		createUser: createUser,
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

	if err != nil {
		errResponse := parsers.HttpErrorParser(errBusiness, ctx, nil)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	render.Status(r, 200)
	render.JSON(rw, r, map[string]interface{}{"data": *result})
}

func (c *userController) Get(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "request_id", uuid.New().String())

	var err error
	var query dtos.QueryParams

	query.BuildQuery(r.URL.Query())

	if err != nil {
		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

}
