package utils

import (
	"errors"
	"fmt"
	"time"

	"edc-proxy/config"
	"io"
	"net/http"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func fetchKeycloakJWKSet(jwksURL string) (jwk.Set, error) {
	resp, err := http.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKs: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch JWKs: status %d", resp.StatusCode)
	}

	// Read the JWK Set
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read JWKs response body: %v", err)
	}

	// Parse the JWK Set
	keySet, err := jwk.Parse(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWKs: %v", err)
	}

	return keySet, nil
}

func validateJWT(tokenStr string, keySet jwk.Set) (jwt.Token, error) {
	issuer := config.GetTrustedIssuer()

	// fmt.Println("!!!! Disable WithAcceptableSkew before building for production!!!!")
	// Parse and validate the JWT token
	token, err := jwt.ParseString(
		tokenStr,
		jwt.WithKeySet(keySet), // Use the JWK set for validation
		jwt.WithIssuer(issuer), // ensure the token was issued by omega-x realm
		jwt.WithAudience("account"),
		jwt.WithClaimValue("azp", "omega-x-marketplace"),
		// jwt.WithAudience("omega-x-marketplace"),
		jwt.WithRequiredClaim("sub"),
		jwt.WithRequiredClaim("organization"),
		jwt.WithAcceptableSkew(5*time.Minute),
		jwt.WithValidate(true), // Validate token claims like `exp`, `iat`
		// jwt.WithHeaderKey("Authorization"), // only works if HTTP request is available in the context
	)
	if err != nil {
		return nil, fmt.Errorf("failed to validate JWT: %v", err)
	}

	return token, nil
}

func IntrospectJWT(tokenString string) (jwt.Token, error) {
	jwksURL := config.GetJWKSURL()

	// Fetch JWKs from Keycloak
	keySet, err := fetchKeycloakJWKSet(jwksURL)
	if err != nil {
		errorMessage := fmt.Sprintf("Error fetching Keycloak JWKs: %v", err)
		return nil, errors.New(errorMessage)
	}

	// Validate the JWT
	token, err := validateJWT(tokenString, keySet)
	if err != nil {
		errorMessage := fmt.Sprintf("Error validating token: %v", err)
		return nil, errors.New(errorMessage)
	}

	return token, nil
}

func GetTokenClaim(token jwt.Token, claimKey string) (string, error) {
	// Access specific claim
	claim, ok := token.Get(claimKey)
	if !ok {
		return "", errors.New("subject (sub) claim not found in token")
	}
	// Convert to strings
	claimString, _ := claim.(string)
	return claimString, nil
}
