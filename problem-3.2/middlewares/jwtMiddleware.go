package middlewares

import (
	"time"
	"github.com/golang-jwt/jwt"
)

func CreateToken(id int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("jwt-super-ultra-secure-?-maybe-yes-maybe-no"))
}