package auth

import (
	"net/http"
	"fmt"
	"../models"
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../env"
  "golang.org/x/crypto/bcrypt"
  "github.com/dgrijalva/jwt-go"
  "time"
)

type LoginResponse struct {
  Token string
  UserName string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

  ((*(&w))).Header().Set("Access-Control-Allow-Origin", "*")
  fmt.Println("here!")

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

  // generate token
  claims := jwt.MapClaims{}
  claims["authorized"] = true
  claims["user_id"] = user.UserName
  claims["exp"] = time.Now().Add(time.Minute * (24*60)).Unix()

  tokenGeneratorWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  token, err := tokenGeneratorWithClaims.SignedString([]byte("secret"))
  if err != nil {
    http.Error(w, err.Error(), http.StatusForbidden)
  }

  var response LoginResponse;
  response.UserName = user.UserName
  response.Token = token

  jResponse, err := json.Marshal(response)
  if err != nil {
    http.Error(w, err.Error(), http.StatusForbidden)
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(jResponse)
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
