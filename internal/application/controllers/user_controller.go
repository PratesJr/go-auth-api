package controllers

import "go-auth-api/internal/domain/adapters"

type userController struct{}

func UserController() adapters.UsersController {
	return &userController{}
}
