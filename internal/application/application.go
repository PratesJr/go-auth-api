package application

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	userRouter "go-auth-api/internal/application/router"
	"go-auth-api/internal/domain/adapters"
)

func Application(userController adapters.UsersController) *chi.Mux {
	router := chi.NewRouter()
	fmt.Println("Server Started...")

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.Recoverer,
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		}),
	)
	router.Route("/auth-api", func(r chi.Router) {
		r.Mount("/user/v1", userRouter.SetupUserRoutes(userController, nil))
	})

	return router
}
