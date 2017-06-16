package models

import (
	"log"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/astaxie/beego"
	"crypto/md5"
	"io"
	"strconv"
	"fmt"
)

func init() {
}

func AddToken(u User, d string) string {
	fmt.Println(u.Username)
	// user id
	var uid int = 0
	// current timestamp
	currentTimestamp := time.Now().UTC().Unix()
	var ttl int64 = 3600
	// md5 of sub & iat
	h := md5.New()
	io.WriteString(h, strconv.Itoa(uid))
	io.WriteString(h, strconv.FormatInt(int64(currentTimestamp), 10))
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": uid,
		"iat": currentTimestamp,
		"exp": currentTimestamp + ttl,
		"nbf": currentTimestamp,
		"iss": d,
		"jti": h.Sum(nil),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(beego.AppConfig.String("HMACKEY")))

	if err != nil {
    	log.Fatal(err)
	}

	return (tokenString)
}
