package auth

import (
	"Ecommerce-Go/config"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateJWT(userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSecond)
	secret := config.Envs.JWTSecret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
