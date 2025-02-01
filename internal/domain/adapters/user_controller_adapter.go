package adapters

import (
	"net/http"
)

type UsersController interface {
	Post(rw http.ResponseWriter, r *http.Request)
	Put(rw http.ResponseWriter, r *http.Request)
	List(rw http.ResponseWriter, r *http.Request)
	FindById(rw http.ResponseWriter, r *http.Request)
}
