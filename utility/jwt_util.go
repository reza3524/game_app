package utility

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const signKey = "secret"

type Claims struct {
	jwt.RegisteredClaims
	UserId uint `json:"userId"`
}

func GenerateToken(userId uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		UserId: userId,
	}
	return token.SignedString([]byte(signKey))
}
