package edcclient

import (
	"fmt"
)

func (c *APIClient) StartContractNegotiation(contractNegotiation AnyJSON) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/contractnegotiations", c.config.ManagementURL)
	return c.makeRequest("POST", url, contractNegotiation)
}

func (c *APIClient) GetContractNegotiations(requestPayload QueryPayload) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/contractnegotiations/request", c.config.ManagementURL)

	if requestPayload == (QueryPayload{}) {
		requestPayload = QueryPayload{
			Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
			Offset: 0,
			Limit:  20,
		}
	}

	return c.makeRequest("POST", url, requestPayload)
}

func (c *APIClient) RetrieveContractNegotiation(contractNegotiationID string) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/contractnegotiations/%s", c.config.ManagementURL, contractNegotiationID)
	return c.makeRequest("GET", url, nil)
}

func (c *APIClient) CancelContractNegotiation(contractNegotiationID string) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/contractnegotiations/%s/cancel", c.config.ManagementURL, contractNegotiationID)
	return c.makeRequest("POST", url, nil)
}

func (c *APIClient) DeclineContractNegotiation(contractNegotiationID string) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/contractnegotiations/%s/decline", c.config.ManagementURL, contractNegotiationID)
	return c.makeRequest("POST", url, nil)
}
