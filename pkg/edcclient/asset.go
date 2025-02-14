package edcclient

import (
	"fmt"
)

func (c *APIClient) CreateAsset(asset AnyJSON) ([]byte, error) {
	url := fmt.Sprintf("%s/v3/assets", c.config.ManagementURL)
	return c.makeRequest("POST", url, asset)
}

func (c *APIClient) UpdateAsset(asset AnyJSON) ([]byte, error) {
	url := fmt.Sprintf("%s/v3/assets", c.config.ManagementURL)
	return c.makeRequest("PUT", url, asset)
}

func (c *APIClient) GetAssets(requestPayload QueryPayload) ([]byte, error) {
	url := fmt.Sprintf("%s/v3/assets/request", c.config.ManagementURL)

	if requestPayload == (QueryPayload{}) {
		requestPayload = QueryPayload{
			Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
			Offset: 0,
			Limit:  20,
		}
	}

	return c.makeRequest("POST", url, requestPayload)
}

func (c *APIClient) DeleteAsset(assetID string) ([]byte, error) {
	url := fmt.Sprintf("%s/v3/assets/%s", c.config.ManagementURL, assetID)
	return c.makeRequest("DELETE", url, nil)
}
