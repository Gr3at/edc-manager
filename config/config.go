package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
}

func GetHost() string {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost" // Default host if not provided
	}
	return host
}

func GetPort() string {
	// Fetch the port from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not provided
	}
	return port
}

func GetSecret() string {
	// Fetch the secret from the environment variable
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "default-server-secret" // Default secret if not provided
	}
	return secret
}

func GetJWKSURL() string {
	// Fetch the jwksUrl from the environment variable
	jwksUrl := os.Getenv("JWKS_URL")
	if jwksUrl == "" {
		jwksUrl = "https://daps.ds.energy.tecnalia.dev/realms/omega-x/protocol/openid-connect/certs" // Default jwksUrl if not provided
	}
	return jwksUrl
}

func GetTrustedIssuer() string {
	// Fetch the issuer from the environment variable
	issuer := os.Getenv("ISSUER")
	if issuer == "" {
		issuer = "https://daps.ds.energy.tecnalia.dev/realms/omega-x" // Default issuer if not provided
	}
	return issuer
}
