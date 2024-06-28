package controllers

import (
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
	ctx := r.Context()
	var payload dtos.UsersDto
	var err error

	err = render.DecodeJSON(r.Body, &payload)

	if err != nil {
		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

	err = validators.Validate(payload)

	if err != nil {
		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

	err, result := c.useCase.Create(ctx, &payload)

	if err != nil {
		render.Status(r, 500)
		render.JSON(rw, r, map[string]string{})

		return
	}

	if err != nil {
		render.Status(r, 500)
		render.JSON(rw, r, map[string]string{})

		return
	}
	render.Status(r, 201)

	render.JSON(rw, r, result)
}

func (c *userController) UpdateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var payload dtos.UpdateUserDto
	var err error

	err = render.DecodeJSON(r.Body, &payload)

	if err != nil {
		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

	err = validators.Validate(payload)

	if err != nil {
		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

	id := chi.URLParam(r, "id")

	parsedUuid, err := uuid.Parse(id)

	if err != nil {
		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

	err, result := c.useCase.Update(ctx, &payload, parsedUuid)

	if err != nil {
		render.Status(r, 500)
		render.JSON(rw, r, map[string]string{})
		return
	}

	render.Status(r, 200)
	render.JSON(rw, r, result)
}
func (c *userController) FindUser(rw http.ResponseWriter, r *http.Request) {
	var err error
	var query dtos.QueryParams

	query.BuildQuery(r.URL.Query())

	err = validators.Validate(query)

	if err != nil {
		render.Status(r, 400)
		render.JSON(rw, r, map[string]string{})

		return
	}

}
