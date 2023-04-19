package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

var SigningKey = []byte(viper.GetString("jwt.signingKey"))

type JwtCustomClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(id int, name string) (string, error) {
	MyJwtCustomClaims := JwtCustomClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((viper.GetDuration("jwt.tokenExpire")) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJwtCustomClaims)

	return token.SignedString(SigningKey)
}

func ParseToken(tokenString string) (JwtCustomClaims, error) {
	MyJwtCustomClaims := JwtCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, &MyJwtCustomClaims, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})

	if err == nil && !token.Valid {
		err = errors.New("invalid Token")
	}
	return MyJwtCustomClaims, err
}

func IsTokenValid(tokenString string) bool {
	_, err := ParseToken(tokenString)
	if err != nil {
		return false
	}
	return true
}
