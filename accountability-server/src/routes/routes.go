package routes

import (
	"net/http"
	"github.com/gorilla/mux"
)


func Handlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/").Methods("GET")

	return r
}