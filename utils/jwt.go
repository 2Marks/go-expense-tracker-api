package utils

import (
	"fmt"
	"time"

	"github.com/2marks/go-expense-tracker-api/config"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(config.Envs.JwtSecret)

func GenerateAuthToken(userId int) (string, error) {
	expirationTime := time.Duration(config.Envs.JwtExpirationInHours) * time.Hour
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"expiresAt": time.Now().Add(expirationTime).Unix(),
		"issuedAt":  time.Now().Unix(),
	})

	signedString, err := claims.SignedString(secretKey)

	if err != nil {
		fmt.Printf("erorr while generating auth token. err:%s", err.Error())
		return "", fmt.Errorf("error occured while generating token")
	}

	return signedString, nil
}

func ValidateAuthToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token sent")
	}

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	return userId.(int), nil
}
