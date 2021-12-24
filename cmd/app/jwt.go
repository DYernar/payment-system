package main

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func (app *application) parseToken(token string, isAccess bool) (string, error) {
	JWTToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("failed to extract token metadata, unexpected signing method: %v", token.Header["alg"])
		}
		if isAccess {
			return []byte(app.config.Server.AccessSecret), nil
		}
		return []byte(app.config.Server.RefreshSecret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := JWTToken.Claims.(jwt.MapClaims)

	var login string

	if ok && JWTToken.Valid {
		login, ok = claims["login"].(string)
		if !ok {
			return "", fmt.Errorf("field id not found")
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			return "", fmt.Errorf("field exp not found")
		}

		expiredTime := time.Unix(int64(exp), 0)
		log.Printf("Expired: %v", expiredTime)
		if time.Now().After(expiredTime) {
			return "", fmt.Errorf("token expired")
		}
		return login, nil
	}

	return "", fmt.Errorf("invalid token")
}
