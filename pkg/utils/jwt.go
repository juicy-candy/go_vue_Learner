package utils

import (
	"ginvue/pkg/model"
	"time"

	"github.com/golang-jwt/jwt"
)

var TOKEN_KEY = []byte("THIS IS A TOKEN KEY")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func GetToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	clams := Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "localhost",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)
	tokenString, err := token.SignedString([]byte(TOKEN_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return TOKEN_KEY, nil
	})
	return token, claims, err
}
