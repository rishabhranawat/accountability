package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../env"
	authmiddleware "../middleware"
	"../models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	Token        string
	RefreshToken string
	UserName     string
}

type CreateUserResponse struct {
	UserName string
}

type LogoutResponse struct {
	ok string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var response LoginResponse

	if authmiddleware.Validate(w, r) {
		cookie, cookieFetchError := r.Cookie("AuthToken")
		if cookieFetchError == nil {
			authToken := cookie.Value
			claims, worked := authmiddleware.ExtractClaims(authToken)
			if worked {
				userID := claims["user_id"]

				if str, ok := userID.(string); ok {
					response.UserName = str
					jResponse, err := json.Marshal(response)
					if err != nil {
						http.Error(w, err.Error(), http.StatusForbidden)
					}
					w.Header().Set("Content-Type", "application/json")
					w.Write(jResponse)
					return
				}

			}
		}
	}

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
	authmiddleware.GenerateTokensAndSetOnHeader(user.UserName, &w)
	response.UserName = user.UserName

	jResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here!")
	authmiddleware.ExpireTokenOnHeader(&w)

	var response LogoutResponse
	response.ok = "ok"

	jResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)

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

	authmiddleware.GenerateTokensAndSetOnHeader(p.UserName, &w)

	var response CreateUserResponse
	response.UserName = p.UserName

	jResponse, err := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jResponse)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// authCookie, _ := r.Cookie("AuthToken")
	// claims := authmiddleware.GetClaims(authCookie.Value)

	// fmt.Println(claims)
	// claimsD, _ := claims.(map[string]interface{})
	// var user models.User
	// fmt.Println(claimsD)
	// fmt.Println(claimsD["user_id"])
	// env.DbConnection.First(&user, "user_name = ?", claimsD["user_id"])
	// if &user == nil {
	// 	http.Error(w, "Cannot find user", http.StatusForbidden)
	// 	return
	// }

	// fmt.Println(user)

	// jResponse, err := json.Marshal(user)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusForbidden)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(jResponse)
}
