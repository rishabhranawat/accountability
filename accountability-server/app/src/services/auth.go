package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../env"
	"../models"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Token        string
	RefreshToken string
	UserName     string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookie("AuthToken"))
	var p models.User
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User
	env.DbConnection.First(&user, "Email = ?", p.Email, p.UserName)
	if &user == nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if error != nil {
		http.Error(w, error.Error(), http.StatusForbidden)
		return
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.UserName
	claims["exp"] = time.Now().Add(time.Minute * (2 * 60)).Unix()

	tokenGeneratorWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenGeneratorWithClaims.SignedString([]byte("secret"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	refreshClaims := jwt.MapClaims{}
	refreshClaims["user_id"] = user.UserName
	refreshClaims["exp"] = time.Now().Add(time.Minute * (24 * 60)).Unix()

	refreshTokenGeneratorWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshTokenGeneratorWithClaims.SignedString([]byte("secret"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	generateTokensAndSetOnHeader(&w, token, refreshToken)
	var response LoginResponse
	response.UserName = user.UserName
	response.RefreshToken = refreshToken
	response.Token = token

	jResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Cookie("AuthToken"))
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
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

func generateTokensAndSetOnHeader(w *http.ResponseWriter, token string, refreshToken string) {

	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(*w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(*w, &refreshCookie)
}
