package auth

import (
	"net/http"
	"fmt"
	"../models"
	"encoding/json"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var p models.User

	fmt.Println(&r.Body)
	err := json.Unmarshal(r.Body, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User: %+v", p)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request){

}

func CreateHandler(w http.ResponseWriter, r *http.Request){
	
}
