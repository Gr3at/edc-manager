package edcclient

import (
	"fmt"
)

type ContractDefinition struct {
	ID               string         `json:"@id"`
	Type             string         `json:"@type"`
	AccessPolicyID   string         `json:"https://w3id.org/edc/v0.0.1/ns/accessPolicyId"`
	ContractPolicyID string         `json:"https://w3id.org/edc/v0.0.1/ns/contractPolicyId"`
	AssetsSelector   []CriterionDto `json:"https://w3id.org/edc/v0.0.1/ns/assetsSelector"`
}

type CriterionDto struct {
	Type         string `json:"@type"`
	OperandLeft  string `json:"https://w3id.org/edc/v0.0.1/ns/operandLeft"`
	Operator     string `json:"https://w3id.org/edc/v0.0.1/ns/operator"`
	OperandRight string `json:"https://w3id.org/edc/v0.0.1/ns/operandRight"`
}

func (c *APIClient) CreateContractDefinition(contractDefinition ContractDefinition) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/contractdefinitions", c.config.ManagementURL)
	return c.makeRequest("POST", url, contractDefinition)
}

func (c *APIClient) GetContractDefinitions(requestPayload QueryPayload) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/contractdefinitions/request", c.config.ManagementURL)

	if requestPayload == (QueryPayload{}) {
		requestPayload = QueryPayload{
			Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
			Offset: 0,
			Limit:  20,
		}
	}

	return c.makeRequest("POST", url, requestPayload)
}

func (c *APIClient) DeleteContractDefinition(contractDefinitionID string) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/contractdefinitions/%s", c.config.ManagementURL, contractDefinitionID)
	return c.makeRequest("DELETE", url, nil)
}
