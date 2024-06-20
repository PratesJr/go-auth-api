package controllers

import (
	"go-auth-api/internal/domain/adapters"
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

func (c *userController) NewUser(rw http.ResponseWriter, r *http.Request) {}

func (c *userController) UpdateUser(rw http.ResponseWriter, r *http.Request) {}
func (c *userController) FindUser(rw http.ResponseWriter, r *http.Request)   {}
