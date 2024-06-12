package auth

import (
	"Ecommerce-Go/config"
	"Ecommerce-Go/types"
	"Ecommerce-Go/utils"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

type contextKey string

const UserKey contextKey = "userID"

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

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from the user request
		tokenString := getTokenFromRequest(r)
		// validate the JWT
		token, err := validateToken(tokenString)
		if err != nil {
			log.Println("failed to validate token : %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("invalid token")
			permissionDenied(w)
			return
		}

		// if is we need to fetch the userID from the DB (id from the token)
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(float64)
		userID := int(str)

		u, err := store.GetUserByID(userID)
		if err != nil {
			log.Println("failed to get user : %v", err)
			permissionDenied(w)
			return
		}

		// set context "userID" to user ID
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)
		handlerFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if token != "" {
		return token
	}
	return ""
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil
	})
}
func permissionDenied(w http.ResponseWriter) {
	utils.WirteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return userID
}
