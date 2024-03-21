package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var securityKey = []byte("ssecretpassword")

// GenerateToken generates a JWT token with the user ID as part of the claims
func GenerateToken(userId uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // token valid for 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(securityKey)
}

// VerifyToken verifies a token JWT validate
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the tokenString
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check the singing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid singing method")
		}

		return securityKey, nil
	})

	// Check for errors
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("Invalid token")
}
