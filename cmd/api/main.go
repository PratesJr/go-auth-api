package main

import (
	"fmt"
	"go-auth-api/internal/application"
	"go-auth-api/internal/application/config"
	"net/http"
)

func main() {
	dependencies := config.GetInjector().InjectDependencies()

	r := application.Application(dependencies.UserController)

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		fmt.Printf("error while starting server")
		return
	}

}
