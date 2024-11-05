package edcclient

import (
	"fmt"
)

type TerminateTransfer struct {
	Type   string `json:"@type"`
	Reason string `json:"https://w3id.org/edc/v0.0.1/ns/reason"`
}

func (c *APIClient) StartDataPushTransferProcess(contractNegotiation AnyJSON) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/transferprocesses", c.config.ManagementURL)
	return c.makeRequest("POST", url, contractNegotiation)
}

func (c *APIClient) GetTransferProcesses(requestPayload QueryPayload) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/transferprocesses/request", c.config.ManagementURL)

	if requestPayload == (QueryPayload{}) {
		requestPayload = QueryPayload{
			Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
			Offset: 0,
			Limit:  20,
		}
	}

	return c.makeRequest("POST", url, requestPayload)
}

func (c *APIClient) TerminateTransferProcess(transferProcessID string, reasonOfTermination string) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/transferprocesses/%s/terminate", c.config.ManagementURL, transferProcessID)

	requestPayload := TerminateTransfer{
		Type:   "https://w3id.org/edc/v0.0.1/ns/TerminateTransfer",
		Reason: reasonOfTermination,
	}

	return c.makeRequest("POST", url, requestPayload)
}
