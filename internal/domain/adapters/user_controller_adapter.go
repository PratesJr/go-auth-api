package adapters

import (
	"net/http"
)

type UsersController interface {
	NewUser(rw http.ResponseWriter, r *http.Request)
	UpdateUser(rw http.ResponseWriter, r *http.Request)
	FindUser(rw http.ResponseWriter, r *http.Request)
}
