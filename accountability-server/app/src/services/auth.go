package auth

import (
	"net/http"
	"fmt"
	"../models"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../env"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var p models.User
	
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	env.DbConnection.First(&user, "email = ?", p.Email)

	fmt.Fprintf(w, "Found User: %+v", user.UserName)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request){

}

func CreateHandler(w http.ResponseWriter, r *http.Request){
	var p models.User

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	env.DbConnection.Create(&p)
	

	fmt.Fprintf(w, "User: %+v", p)
}