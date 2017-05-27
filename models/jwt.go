package models

import (
	"log"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/astaxie/beego"
)

func init() {
}

func AddToken(u User) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": u.Username,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(beego.AppConfig.String("HMACKEY")))

	if err != nil {
    	log.Fatal(err)
	}

	return (tokenString)
}
