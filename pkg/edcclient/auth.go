package edcclient

import (
	"bytes"
	"edc-proxy/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// AuthStrategy defines the interface for different authentication strategies.
type AuthStrategy interface {
	// Authenticate(req *http.Request)
	SetAuthHeader(req *http.Request)
}

type AuthStrategyConfig struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
	APIKey       string
}

func NewAuthStrategy(authType string, sConf AuthStrategyConfig) AuthStrategy {
	var authStrategy AuthStrategy
	if authType == "client_credentials" {
		authStrategy = &JWTAuth{
			ClientID:     sConf.ClientID,
			ClientSecret: sConf.ClientSecret,
			TokenURL:     sConf.TokenURL,
		}
	} else {
		authStrategy = &APIKeyAuth{
			APIKey: sConf.APIKey,
		}
	}
	return authStrategy
}

// JWTAuth is a concrete strategy that handles JWT Bearer Token authentication.
type JWTAuth struct {
	ClientID     string // For OAuth2 client credentials
	ClientSecret string
	TokenURL     string
	accessToken  string
}

// SetAuthHeader adds the JWT Bearer Token to the request.
// Before setting the bearer access token, it retrieves a fresh token from the token url
func (j *JWTAuth) SetAuthHeader(req *http.Request) {
	token, err := j.getOAuthToken()
	if (err != nil) || (token == "") {
		utils.Log.Errorf("Error retrieving access token from %s: %v", j.TokenURL, err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+j.accessToken)
}

// getOAuthToken retrieves an access token using client credentials.
func (j *JWTAuth) getOAuthToken() (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", j.ClientID)
	data.Set("client_secret", j.ClientSecret)

	http_client := &http.Client{}
	req, err := http.NewRequest("POST", j.TokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http_client.Do(req)
	if err != nil {
		fmt.Printf("Error getting auth token: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get OAuth token")
	}

	var tokenResp struct {
		AccessToken      string `json:"access_token"`
		ExpiresIn        int    `json:"expires_in"`
		RefreshExpiredIn int    `json:"refresh_expires_in"`
		TokenType        string `json:"token_type"`
		NotBeforePolicy  int    `json:"not-before-policy"`
		Scope            string `json:"scope"`
	}

	err = json.NewDecoder(resp.Body).Decode(&tokenResp)
	if err != nil {
		return "", err
	}

	j.accessToken = tokenResp.AccessToken
	return tokenResp.AccessToken, nil
}

// APIKeyAuth is a concrete strategy that handles X-API-Key authentication.
type APIKeyAuth struct {
	APIKey string
}

// SetAuthHeader adds the X-API-Key to the request.
func (a *APIKeyAuth) SetAuthHeader(req *http.Request) {
	req.Header.Add("X-Api-Key", a.APIKey)
}
