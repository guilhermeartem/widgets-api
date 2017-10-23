package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//constants for testing purposes
const (
	ValidUser = "username"
	ValidPass = "password"
	SecretKey = "AReallyDificultSecretKeyShouldBeHere"
)

//UserCredentials struct for login
type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Token struct for authentication response
type Token struct {
	Token string `json:"token"`
}

//LoginHandler handler for user login`
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user UserCredentials

	err := json.NewDecoder(r.Body).Decode(&user)

	log.Println(user)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	if user.Username != ValidUser || user.Password != ValidPass {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		fmt.Fprint(w, "Invalid credentials")
		return
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		log.Fatal(err)
	}

	response := Token{tokenString}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}

}

//ValidateTokenMiddleware middleware for validating if user is authenticated in the system
func ValidateTokenMiddleware(inner http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("autenticando")
		token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (interface{}, error) {
				return []byte(SecretKey), nil
			})

		if err == nil {
			if token.Valid {
				inner.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Token is not valid")
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			log.Println(err)
			fmt.Fprint(w, "Unauthorized access to this resource")
		}
	})

}
