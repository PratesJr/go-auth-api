package main

import (
	"fmt"
	"go-auth-api/internal/application/router"
	"net/http"
)

func main() {
	r := router.RegisterRoutes()

	err := http.ListenAndServe(":3000", r)

	if err != nil {
		_ = fmt.Errorf("deu ruim")
		return
	}

}
