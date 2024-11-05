package edcclient

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type AnyJSON map[string]interface{} // is a bit less generic than interface{} to be able to support unknown JSON input

func printJSONOrString(byteData []byte) {
	var prettyJSON []byte
	var jsonData interface{}

	if err := json.Unmarshal(byteData, &jsonData); err == nil {
		// Use MarshalIndent to pretty-print with indentation
		if prettyJSON, err = json.MarshalIndent(jsonData, "", "  "); err == nil {
			fmt.Printf("received data (json): %s\n", prettyJSON)
		} else {
			fmt.Printf("Error formatting JSON: %v\n", err)
		}
	} else {
		fmt.Printf("received data (safe string): %s\n", string(byteData))
	}
}

func waitForUserInput(prompt string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	_, _ = reader.ReadString('\n') // Waits for the user to press Enter
}

// QueryPayload defines the expected schema to query assets, policies, contract definitions
type QueryPayload struct {
	Type   string `json:"@type"`
	Offset int    `json:"https://w3id.org/edc/v0.0.1/ns/offset"`
	Limit  int    `json:"https://w3id.org/edc/v0.0.1/ns/limit"`
}
