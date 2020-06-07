package routes

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"../services"
)


func Handlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", AccountabilityAppHandler).Methods("GET")

	// auth services
	r.HandleFunc("/auth/create", auth.CreateHandler).Methods("POST")
	r.HandleFunc("/auth/login", auth.LoginHandler).Methods("POST")
	r.HandleFunc("/auth/logout", auth.LogoutHandler).Methods("POST")


	return r
}

func AccountabilityAppHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Accountability Server is up")
}