package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserId    string `json:"user_id"`
	SessionId string `json:"session_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userId string, sessionId string) (string, error) {
	expirationTime := time.Now().Add(6 * time.Hour)

	claims := &MyCustomClaims{
		UserId:    userId,
		SessionId: sessionId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	envSecret := os.Getenv("SECRET_JWT")

	tokenString, err := token.SignedString([]byte(envSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
