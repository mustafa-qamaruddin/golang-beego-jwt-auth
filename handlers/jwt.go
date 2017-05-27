package handlers

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/astaxie/beego"
)

func init() {
}

func Jwt(ctx *context.Context) {
		var uri string = ctx.Input.URI()
		if  uri == "/v1/jwt" {
			return
		}

		if ctx.Input.Header("Authorization") == "" {
			ctx.Output.SetStatus(403)
			ctx.Output.Body([]byte("notAllowed"))
		}

		if !verify(ctx.Input.Header("Authorization")) {
			ctx.Output.SetStatus(403)
			ctx.Output.Body([]byte(ctx.Input.Header("Authorization")))
		}
}

func verify(tokenString string) bool {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(beego.AppConfig.String("HMACKEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["uid"], claims["nbf"])
		return true
	} else {
		fmt.Println(err)		
		return false
	}
}