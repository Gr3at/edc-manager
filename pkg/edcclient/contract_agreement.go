package edcclient

import (
	"fmt"
)

func (c *APIClient) GetContractAgreements(requestPayload QueryPayload) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/contractagreements/request", c.config.ManagementURL)

	if requestPayload == (QueryPayload{}) {
		requestPayload = QueryPayload{
			Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
			Offset: 0,
			Limit:  20,
		}
	}

	return c.makeRequest("POST", url, requestPayload)
}
