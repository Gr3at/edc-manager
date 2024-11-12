package services

import (
	"edc-proxy/config"
	"edc-proxy/models"
	"edc-proxy/pkg/edcclient"
	"edc-proxy/utils"
	"fmt"
	"strings"
)

// var CreateResources = func(data map[string]interface{}) error {
// 	// Prepare the request to third-party service
// 	jsonData, _ := json.Marshal(data)
// 	resp, err := http.Post("https://third-party-service.com/api/resource", "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Handle the response (log or process further if needed)
// 	if resp.StatusCode != http.StatusCreated {
// 		return err
// 	}

// 	return nil
// }

func SetupAPIClient(orgID string) (*edcclient.APIClient, error) {
	// 1. retrieve connector from db data
	var orgConnector models.Connector
	result := config.DB.Where("org_id = ?", orgID).First(&orgConnector)

	if result.RowsAffected == 0 {
		utils.Log.Warnf("No connector credentials found for the requested org: %s", orgID)
		return nil, fmt.Errorf("you need to link your organization (%s) Connector first", orgID)
	}

	if result.Error != nil {
		utils.Log.Errorf("Error fetching the organization's connector: %v", result.Error)
		return nil, fmt.Errorf("unable to use your connector. You need to link your organization (%s) Connector first", orgID)
	}

	// 2. setup and link connector
	secretKey := config.GetSecret()
	decryptedPayload, err := utils.DecryptKey(orgConnector.Credentials, secretKey)
	if err != nil {
		utils.Log.Errorf("Error decrypting connector credentials for orgID %s. error details: (%v)", orgID, err)
		return nil, fmt.Errorf("unable to communicate with your connector. Did you provide correct credentials?")
	}

	// if apiKey then one string slice, else if clientCredentials then 2 string slices
	var aSConf edcclient.AuthStrategyConfig
	if string(orgConnector.CredentialsType) == "client_credentials" {
		decryptedCredentials := strings.Split(decryptedPayload, ":")
		if len(decryptedCredentials) != 2 {
			utils.Log.Errorf("Error unwrapping connector credentials for orgID %s. error details: (%v)", orgID, err)
			return nil, fmt.Errorf("unable to communicate with your connector. Are the provided credentials correct")
		}

		clientID := decryptedCredentials[0]
		clientSecret := decryptedCredentials[1]
		aSConf = edcclient.AuthStrategyConfig{ClientID: clientID, ClientSecret: clientSecret, TokenURL: "https://keycloak.prod-sovity.azure.sovity.io/realms/Portal/protocol/openid-connect/token"}
	} else if string(orgConnector.CredentialsType) == "api_key" {
		aSConf = edcclient.AuthStrategyConfig{APIKey: decryptedPayload}
	} else {
		utils.Log.Errorf("Not a supported credential type for orgID %s.", orgID)
		return nil, fmt.Errorf("unsupported EDC authorization method (%s)", string(orgConnector.CredentialsType))
	}
	authStrategy := edcclient.NewAuthStrategy(string(orgConnector.CredentialsType), aSConf)

	config := edcclient.Config{
		ManagementURL: orgConnector.APIUrl,
		AuthStrategy:  authStrategy,
	}

	utils.Log.Info("Get a new API Client instance")
	apiClient, err := edcclient.NewAPIClient(config, nil)

	if err != nil {
		utils.Log.Errorf("error while creating API client from factory. error details: (%v)", err)
		return nil, fmt.Errorf("error while creating EDC API client")
	}
	return apiClient, nil
}
