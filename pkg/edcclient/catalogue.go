package edcclient

import (
	"fmt"
)

type CatalogueRequestBody struct {
	Type                string       `json:"@type"`
	Protocol            string       `json:"https://w3id.org/edc/v0.0.1/ns/protocol"`
	CounterPartyAddress string       `json:"https://w3id.org/edc/v0.0.1/ns/counterPartyAddress"`
	QuerySpec           QueryPayload `json:"https://w3id.org/edc/v0.0.1/ns/querySpec"`
}

func (c *APIClient) RequestCatalogue(requestPayload CatalogueRequestBody) ([]byte, error) {
	url := fmt.Sprintf("%s/v2/catalog/request", c.config.ManagementURL)

	if requestPayload == (CatalogueRequestBody{}) {
		return []byte{}, fmt.Errorf("you need to provide a proper request body to retrieve EDC catalogue data")
		// requestPayload = CatalogueRequestBody{
		// 	Type:                "https://w3id.org/edc/v0.0.1/ns/CatalogRequest",
		// 	Protocol:            "dataspace-protocol-http",
		// 	CounterPartyAddress: "{{PROVIDER_EDC_PROTOCOL_URL}}",
		// 	QuerySpec: QueryPayload{
		// 		Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
		// 		Offset: 0,
		// 		Limit:  10,
		// 	},
		// }
	}

	return c.makeRequest("POST", url, requestPayload)
}
