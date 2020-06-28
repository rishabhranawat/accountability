package routes

import (
	"fmt"
	"net/http"

	auth "../services"
	"../taskhandler"
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", AccountabilityAppHandler).Methods("GET")

	// auth services
	r.HandleFunc("/auth/create", auth.CreateHandler).Methods("POST")
	r.HandleFunc("/auth/login", auth.LoginHandler).Methods("POST")
	r.HandleFunc("/auth/logout", auth.LogoutHandler).Methods("POST")

	// task management services
	r.HandleFunc("/tasks/create-task", taskhandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/update-task", taskhandler.UpdateTask).Methods("POST")
	r.HandleFunc("/tasks/remove-task", taskhandler.RemoveTask).Methods("POST")
	r.HandleFunc("/tasks/get-tasks", taskhandler.GetUserTasks).Methods("GET")

	return r
}

func AccountabilityAppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookie("AuthToken"))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Accountability Server is up")
}
