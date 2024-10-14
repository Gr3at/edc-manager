package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// Secret key used for signing JWT tokens
var jwtSecret = []byte("your-secret-key")

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	// Parse and validate the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func GenerateJWT(userID string) (string, error) {
	// Define the expiration time
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the claims
	claims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: expirationTime.Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	return token.SignedString(jwtSecret)
}
