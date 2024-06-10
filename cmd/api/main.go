package main

import (
	"go-auth-api/internal/application/config"
	"go-auth-api/internal/application/router"
	"net/http"
)

func main() {
	dependencies := config.GetInjector()
	r := router.RegisterRoutes()

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		return
	}

}
