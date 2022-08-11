package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "test-client"
	claims["exp"] = time.Now().Add(time.Minute).Unix()

	tokenString, err := token.SignedString([]byte(GetKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetKey() string {
	return ""
}
