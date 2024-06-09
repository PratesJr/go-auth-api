package main

import (
	"go-auth-api/cmd/router"
	"net/http"
)

func main() {
	r := router.RegisterRoutes()

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		return
	}

}
