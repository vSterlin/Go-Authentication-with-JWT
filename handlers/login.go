package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/vSterlin/jwt-auth/data"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type res struct {
	Token string `json:"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	data.FromJSON(&creds, r.Body)

	if creds.Email != "a@a" || creds.Password != "pass" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = creds.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Errorf("Signing went wrong")
	}
	response := &res{
		Token: t,
	}
	data.ToJSON(response, w)
}
