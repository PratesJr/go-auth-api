package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"go-auth-api/internal/application/validators"
	"go-auth-api/internal/domain/adapters"
	"go-auth-api/internal/domain/dtos"
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

	var payload dtos.UsersDto
	var err error

	err = render.DecodeJSON(r.Body, &payload)

	if err != nil {
		render.Status(r, 400)
		return
	}

	err = validators.Validate(payload)

	if err != nil {
		render.Status(r, 400)
		return
	}

	err, result := c.useCase.Create(&payload)

	if err != nil {
		render.Status(r, 500)
		return
	}

	response, err := json.Marshal(result)

	if err != nil {
		render.Status(r, 500)
	}

	render.JSON(rw, r, response)
}

func (c *userController) UpdateUser(rw http.ResponseWriter, r *http.Request) {

	var payload dtos.UpdateUserDto
	var err error

	err = render.DecodeJSON(r.Body, &payload)

	if err != nil {
		render.Status(r, 400)
		return
	}

	err = validators.Validate(payload)

	if err != nil {
		render.Status(r, 400)
		return
	}

	id := chi.URLParam(r, "id")

	parsedUuid, err := uuid.Parse(id)

	if err != nil {
		render.Status(r, 400)
		return
	}

	err, result := c.useCase.Update(&payload, parsedUuid)

	if err != nil {
		render.Status(r, 500)
		return
	}

	response, err := json.Marshal(result)

	render.JSON(rw, r, response)
}
func (c *userController) FindUser(rw http.ResponseWriter, r *http.Request) {}
