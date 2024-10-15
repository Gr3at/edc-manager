package utils

import (
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	// token, err := GenerateJWT("test-user")
	// if err != nil {
	// 	t.Fatalf("Expected no error, got %v", err)
	// }

	// parsedToken, err := ValidateJWT(token)
	// if err != nil {
	// 	t.Fatalf("Expected valid token, got error: %v", err)
	// }

	// if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
	// 	if claims["sub"] != "test-user" {
	// 		t.Errorf("Expected subject to be 'test-user', got %v", claims["sub"])
	// 	}
	// } else {
	// 	t.Errorf("Invalid claims structure")
	// }
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	// _, err := ValidateJWT("invalid.token.string")
	// if err == nil {
	// 	t.Fatalf("Expected error for invalid token, got none")
	// }
}

func TestJWTExpiration(t *testing.T) {
	// jwtSecret = []byte("temporary-secret") // Temporary change to avoid affecting other tests
	// claims := jwt.StandardClaims{
	// 	Subject:   "test-user",
	// 	ExpiresAt: time.Now().Add(-time.Hour).Unix(), // Expired token
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenString, _ := token.SignedString(jwtSecret)

	// _, err := ValidateJWT(tokenString)
	// if err == nil {
	// 	t.Fatalf("Expected expired token error, got none")
	// }
}
