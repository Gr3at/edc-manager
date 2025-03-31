package edcclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
	Logger     *logrus.Logger
}

// ClientFactory initializes a new Client instance with either API Key or OAuth2 config.
func NewAPIClient(config Config, httpClient *http.Client, logger *logrus.Logger) (*APIClient, error) {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: time.Second * 10,
		}
		// httpClient = http.DefaultClient
	}

	client := &APIClient{
		config:     config,
		httpClient: httpClient,
		Logger:     logger,
	}

	return client, nil
}

func (c *APIClient) SetAuthorizationHeader(req *http.Request) {
	c.config.AuthStrategy.SetAuthHeader(req)
}

// makeRequest is a helper to make HTTP requests.
func (c *APIClient) makeRequest(method, url string, payload interface{}) ([]byte, error) {
	requestID := uuid.New().String()

	// Marshal payload to JSON
	body, err := json.Marshal(payload)
	if err != nil {
		if c.Logger != nil {
			c.Logger.WithError(err).Error("Failed to marshal payload")
		}
		return nil, err
	}

	if c.Logger != nil {
		c.Logger.WithFields(logrus.Fields{
			"requestID": requestID,
			"method":    method,
			"url":       url,
			"body":      string(body),
		}).Info("Making request")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		if c.Logger != nil {
			c.Logger.WithFields(logrus.Fields{
				"requestID": requestID,
			}).WithError(err).Error("Failed to create request")
		}
		return nil, err
	}

	c.SetAuthorizationHeader(req)
	req.Header.Add("Content-Type", "application/json")

	// Execute the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		if c.Logger != nil {
			c.Logger.WithFields(logrus.Fields{
				"requestID": requestID,
			}).WithError(err).Error("Request failed")
		}
		return nil, err
	}
	defer res.Body.Close()

	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		if c.Logger != nil {
			c.Logger.WithFields(logrus.Fields{
				"requestID": requestID,
			}).WithError(err).Error("Failed to read response body")
		}
		return nil, err
	}

	if c.Logger != nil {
		c.Logger.WithFields(logrus.Fields{
			"requestID": requestID,
			"status":    res.StatusCode,
			"body":      string(respBody),
		}).Info("Received response")
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("error: %s", respBody)
	}
	return respBody, nil
}
