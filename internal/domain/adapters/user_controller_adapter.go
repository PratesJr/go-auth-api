package adapters

import (
	"net/http"
)

type UsersController interface {
	Post(rw http.ResponseWriter, r *http.Request)
	Put(rw http.ResponseWriter, r *http.Request)
	Get(rw http.ResponseWriter, r *http.Request)
}
