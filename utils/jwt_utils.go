package utils

import (
	"errors"
	"fmt"
	"strings"
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
		// jwt.WithAcceptableSkew(5*time.Minute),
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
		// debug: check if the token is indeed invalid
		trustedIssuerURL := config.GetTrustedIssuer()
		tokenIsValid, validationErr := isAccessTokenValid(trustedIssuerURL+"/protocol/openid-connect/userinfo", tokenString)
		if validationErr != nil {
			return nil, fmt.Errorf("error while debugging the validity of the oauth token: %v. validation error: %v", validationErr, err)
		}
		if tokenIsValid {
			return nil, fmt.Errorf("the JWT is valid but was wrongly flagged as invalid. validation error: %v", err)
		}
		return nil, fmt.Errorf("error in JWT validation: %v", err)
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

// Because the introspection endpoint does always return {"active":false} regardless if the token is valid or not, we may utilzie the Get User Info endpoint.
// If the response is 200, then the access token is valid.
func isAccessTokenValid(userInfoUrl, accessToken string) (bool, error) {
	payload := strings.NewReader("")
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("GET", userInfoUrl, payload)

	if err != nil {
		return false, fmt.Errorf("failed to wrap user info request. error: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	res, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to wrap user info request. error: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		return true, nil
	} else {
		return false, fmt.Errorf("the access token is invalid. response code from oauth server: %d", res.StatusCode)
	}
}
