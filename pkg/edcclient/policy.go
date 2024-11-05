package edcclient

import (
	"fmt"
)

func (c *APIClient) CreatePolicy(policy AnyJSON) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/policydefinitions", c.config.ManagementURL)
	return c.makeRequest("POST", url, policy)
}

func (c *APIClient) GetPolicies(requestPayload QueryPayload) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/policydefinitions/request", c.config.ManagementURL)

	if requestPayload == (QueryPayload{}) {
		requestPayload = QueryPayload{
			Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
			Offset: 0,
			Limit:  20,
		}
	}

	return c.makeRequest("POST", url, requestPayload)
}

func (c *APIClient) DeletePolicy(policyID string) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/policydefinitions/%s", c.config.ManagementURL, policyID)
	return c.makeRequest("DELETE", url, nil)
}
