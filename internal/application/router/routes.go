package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RegisterRoutes() *chi.Mux {
	routes := chi.NewRouter()
	routes.Use(middleware.Logger)

	return routes
}
