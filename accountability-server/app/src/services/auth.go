package auth

import (
	"net/http"
	"fmt"
	"../models"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../env"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var p models.User
	
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	env.DbConnection.First(&user, "Email = ?", p.Email)
	if &user == nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
    if error != nil {
		http.Error(w, error.Error(), http.StatusForbidden)
		return
    }
    
	fmt.Fprintf(w, "User: %+v is logged in %s", user.UserName, user.Password)
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

	hash, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.MinCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	p.Password = string(hash)
	env.DbConnection.Create(&p)
	
	fmt.Fprintf(w, "Successfully created user with username: %s and email: %s", p.UserName, p.Email)
}