package edcclient

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("Starting JWTAuth based use case test")

	aSConf := AuthStrategyConfig{ClientID: "", ClientSecret: "", TokenURL: ""}
	authType := "JWTAuth"
	authStrategy := NewAuthStrategy(authType, aSConf)

	config := Config{
		ManagementURL: "https://example.com/api/management",
		AuthStrategy:  authStrategy,
	}

	fmt.Println("Get a new API Client instance")
	apiClient, err := NewAPIClient(config, nil)

	if err != nil {
		fmt.Printf("error while creating API client from factory: %v\n", err)
	}

	fmt.Println("Make the Get Assets request")
	byteData, err := apiClient.GetAssets(QueryPayload{
		Type:   "https://w3id.org/edc/v0.0.1/ns/QuerySpec",
		Offset: 0,
		Limit:  20,
	})

	if err != nil {
		fmt.Printf("Error in GetAssets request: %v\n", err)
		t.Errorf("got status %v but wanted %s", err, "")
		panic(err)
	}

	waitForUserInput("Press enter to view the received response...")
	printJSONOrString(byteData)
}
