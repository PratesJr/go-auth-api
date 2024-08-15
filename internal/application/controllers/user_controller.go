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
	useCase adapters.UserUseCase
}

func UserControllerConstructor(useCase adapters.UserUseCase) adapters.UsersController {
	return &userController{
		useCase: useCase,
	}
}

func (c *userController) NewUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "request_id", uuid.New().String())

	var payload dtos.UsersDto

	err := render.DecodeJSON(r.Body, &payload)

	if err != nil {

		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

	result, errBusiness := c.useCase.Create(ctx, &payload)

	if errBusiness != nil {

		errResponse := parsers.HttpErrorParser(errBusiness, ctx, nil)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	render.Status(r, 201)

	render.JSON(rw, r, map[string]interface{}{"data": *result})
}

func (c *userController) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, "request_id", uuid.New().String())

	var payload dtos.UpdateUserDto

	errDecode := render.DecodeJSON(r.Body, &payload)

	if errDecode != nil {
		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

	id := chi.URLParam(r, "id")

	parsedUuid, errUuid := uuid.Parse(id)

	if errUuid != nil {
		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

	result, err := c.useCase.Update(ctx, &payload, parsedUuid)

	if err != nil {
		errResponse := parsers.HttpErrorParser(err, ctx, nil)

		render.Status(r, errResponse.StatusCode)
		render.JSON(rw, r, map[string]exceptions.HttpException{"error": errResponse})

		return
	}

	render.Status(r, 200)
	render.JSON(rw, r, map[string]interface{}{"data": *result})
}
func (c *userController) FindUser(rw http.ResponseWriter, r *http.Request) {
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
