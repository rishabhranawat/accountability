package routes

import (
	"fmt"
	"net/http"

	authmiddleware "../middleware"
	auth "../services"
	feed "../services/feed"
	relationshipmgmt "../services/relationshipmgmt"
	"../taskhandler"
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()

	// auth services
	r.HandleFunc("/auth/create", auth.CreateHandler).Methods("POST")
	r.HandleFunc("/auth/login", auth.LoginHandler).Methods("POST")

	validatedAuthRoutes := r.PathPrefix("/auth").Subrouter()
	validatedAuthRoutes.Use(AuthMiddleware)
	validatedAuthRoutes.HandleFunc("/", AccountabilityAppHandler).Methods("GET")
	validatedAuthRoutes.HandleFunc("/logout", auth.LogoutHandler).Methods("POST")
	validatedAuthRoutes.HandleFunc("/get-user", auth.GetUserHandler).Methods("GET")

	// task management services
	taskRoutes := r.PathPrefix("/tasks").Subrouter()
	taskRoutes.Use(AuthMiddleware)
	taskRoutes.HandleFunc("/create-task", taskhandler.CreateTask).Methods("POST")
	taskRoutes.HandleFunc("/update-task", taskhandler.UpdateTask).Methods("POST")
	taskRoutes.HandleFunc("/remove-task", taskhandler.RemoveTask).Methods("POST")
	taskRoutes.HandleFunc("/fetch-tasks", taskhandler.FetchUserTasks).Methods("POST")
	taskRoutes.HandleFunc("/create-task-comment", taskhandler.PostTaskComment).Methods("POST")
	taskRoutes.HandleFunc("/fetch-task-details/{task-id}", taskhandler.FetchTaskDetails).Methods("GET")
	taskRoutes.HandleFunc("/create-task-update", taskhandler.PostTaskUpdate).Methods("POST")
	taskRoutes.HandleFunc("/fetch-task-comments/{task-id}", taskhandler.FetchTaskComments).Methods("GET")
	taskRoutes.HandleFunc("/fetch-task-updates/{task-id}", taskhandler.FetchTaskUpdates).Methods("GET")

	// feed
	taskRoutes.HandleFunc("/user-feed", feed.GetFeed).Methods("GET")
	taskRoutes.HandleFunc("/user-profile-feed", feed.GetUserSpecificFeed).Methods("GET")


	//relationshipmgmt
	relationshipRoutes := r.PathPrefix("/relationship").Subrouter()
	relationshipRoutes.Use(AuthMiddleware)
	relationshipRoutes.HandleFunc("/create-relationship", relationshipmgmt.CreateRelationship).Methods("POST")
	relationshipRoutes.HandleFunc("/approve-relationship", relationshipmgmt.ApproveRelationship).Methods("POST")
	relationshipRoutes.HandleFunc("/delete-relationship", relationshipmgmt.DeleteRelationship).Methods("POST")

	return r
}

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !authmiddleware.Validate(w, r) {
			http.Error(w, "No valid login token", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func AccountabilityAppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookie("AuthToken"))
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Accountability Server is up")
}
