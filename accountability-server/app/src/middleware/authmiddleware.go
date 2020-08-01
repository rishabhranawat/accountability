package authmiddleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"../env"
	"../models"
	"github.com/dgrijalva/jwt-go"
)

/*
Token generation
*/
func GenerateTokensAndSetOnHeader(UserName string, w *http.ResponseWriter) {
	authToken, errA := generateAuthToken(UserName)
	refreshToken, errR := generateRefreshToken(UserName)

	if errA != nil || errR != nil {
		return
	}
	setTokensOnHeader(w, authToken, refreshToken)
}

func generateAuthToken(UserName string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = UserName
	claims["exp"] = time.Now().Add(time.Minute * (1 * 60)).Unix()

	tokenGeneratorWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenGeneratorWithClaims.SignedString([]byte("secret"))
	if err != nil {
		return token, err
	}
	return token, nil
}

func generateRefreshToken(UserName string) (string, error) {
	refreshClaims := jwt.MapClaims{}
	refreshClaims["user_id"] = UserName
	refreshClaims["exp"] = time.Now().Add(time.Minute * (24 * 60)).Unix()

	refreshTokenGeneratorWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshTokenGeneratorWithClaims.SignedString([]byte("secret"))
	if err != nil {
		return refreshToken, err
	}

	return refreshToken, nil
}

func setTokensOnHeader(w *http.ResponseWriter, token string, refreshToken string) {

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

func ExpireTokenOnHeader(w *http.ResponseWriter) {
	authCookie := http.Cookie{
		Name:    "AuthToken",
		Path:    "/",
		Expires: time.Now(),
	}

	http.SetCookie(*w, &authCookie)

	refreshCookie := http.Cookie{
		Name:    "RefreshToken",
		Path:    "/",
		Expires: time.Now(),
	}

	http.SetCookie(*w, &refreshCookie)
}

/*
Token validation
*/
func Validate(w http.ResponseWriter, r *http.Request) bool {
	return checkAndRefreshToken(w, r)
}

func checkAndRefreshToken(w http.ResponseWriter, r *http.Request) bool {
	isAuthTokenValid, error := isTokenValid(r, "AuthToken")
	if error != nil {
		return false
	}
	if isAuthTokenValid == true {
		return true
	}

	isRefreshTokenValid, error := isTokenValid(r, "RefreshToken")
	if error != nil {
		return false
	}
	if isRefreshTokenValid == true {
		return true
	}

	//Todo: rish, refresh the token
	return false

}

func isTokenValid(r *http.Request, tokenType string) (bool, error) {
	cookie, cookieFetchError := r.Cookie(tokenType)
	if cookieFetchError != nil {
		return false, cookieFetchError
	}
	return validateToken(cookie.Value), nil
}

func validateToken(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// todo move the secret key to an env file
		return []byte("secret"), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		return false
	}
}

// func GetClaims(tokenString string) interface{} {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte("secret"), nil
// 	})

// 	if err != nil {
// 		return nil
// 	}

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims
// 	}
// 	return nil
// }

func GetCurrentUser(r *http.Request) models.User {
	cookie, cookieFetchError := r.Cookie("AuthToken")
	var user models.User
	if cookieFetchError == nil {
		authToken := cookie.Value
		claims, worked := ExtractClaims(authToken)
		if worked {
			maybeUsername := claims["user_id"]

			if str, ok := maybeUsername.(string); ok {
				env.DbConnection.Where("user_name = ?", str).Find(&user)
			}
		}
	}
	return user
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "secret"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
