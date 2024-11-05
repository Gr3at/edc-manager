package edcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Config struct to set up the client with endpoint and credentials.
type Config struct {
	ManagementURL string
	AuthStrategy  AuthStrategy
}

// Client struct holds configuration and HTTP client.
type APIClient struct {
	config     Config
	httpClient *http.Client
}

// ClientFactory initializes a new Client instance with either API Key or OAuth2 config.
func NewAPIClient(config Config, httpClient *http.Client) (*APIClient, error) {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
		// httpClient = http.DefaultClient
	}

	client := &APIClient{
		config:     config,
		httpClient: httpClient,
	}

	return client, nil
}

func (c *APIClient) SetAuthorizationHeader(req *http.Request) {
	c.config.AuthStrategy.SetAuthHeader(req)
}

// makeRequest is a helper to make HTTP requests.
func (c *APIClient) makeRequest(method, url string, payload interface{}) ([]byte, error) {
	// Marshal payload to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	c.SetAuthorizationHeader(req)
	req.Header.Add("Content-Type", "application/json")

	// Execute the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("error: %s", respBody)
	}
	return respBody, nil
}
