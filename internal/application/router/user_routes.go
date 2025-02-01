package router

import (
	"github.com/go-chi/chi/v5"
	"go-auth-api/internal/domain/adapters"
	"net/http"
)

func SetupUserRoutes(controller adapters.UsersController, middlewares []func(http.Handler) http.Handler) *chi.Mux {
	routes := chi.NewRouter()
	routes.Use(middlewares...)

	routes.Post("/user", controller.Post)
	routes.Put("/user/{id}", controller.Put)
	routes.Get("/user", controller.List)

	return routes
}
