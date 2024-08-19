package main

import (
	"fmt"
	"go-auth-api/internal/application"
	"go-auth-api/internal/application/config/dependencies"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	injector := dependencies.GetInjector().InjectDependencies()

	r := application.Application(injector.UserController)

	err := http.ListenAndServe(":3333", r)

	if err != nil {

		fmt.Printf("error while starting server")
		return

	}

	sig := <-sigCh

	fmt.Printf("Caught signal %v:\n", sig)

}
