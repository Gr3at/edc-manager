package edcclient

import (
	"fmt"
)

// Define the request payload structure
type RequestAssetPayload struct {
	Type   string `json:"@type"`
	Offset int    `json:"https://w3id.org/edc/v0.0.1/ns/offset"`
	Limit  int    `json:"https://w3id.org/edc/v0.0.1/ns/limit"`
}

func (c *APIClient) CreateAsset(assetID string, asset AnyJSON) ([]byte, error) {
	url := fmt.Sprintf("%s/v3/assets/%s", c.config.ManagementURL, assetID)
	return c.makeRequest("POST", url, asset)
}

func (c *APIClient) GetAssets(requestAssetPayload AnyJSON) ([]byte, error) {
	url := fmt.Sprintf("%s/v3/assets/request", c.config.ManagementURL)
	var payload AnyJSON
	if requestAssetPayload == nil {
		payload = AnyJSON{
			"@type":                                 "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
			"https://w3id.org/edc/v0.0.1/ns/offset": 0,
			"https://w3id.org/edc/v0.0.1/ns/limit":  20,
		}
	} else {
		payload = requestAssetPayload
	}
	// payload := RequestAssetPayload{
	// 	Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
	// 	Offset: 0,
	// 	Limit:  20,
	// }

	body, err := c.makeRequest("POST", url, payload)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	return body, nil
}
