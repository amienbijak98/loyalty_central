package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWTToken(u interface{}, secretKey string, expiredIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": u,
		"exp":  time.Now().Add(expiredIn).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
