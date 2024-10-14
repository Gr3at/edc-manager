package services

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func CreateResources(data map[string]interface{}) error {
	// Prepare the request to third-party service
	jsonData, _ := json.Marshal(data)
	resp, err := http.Post("https://third-party-service.com/api/resource", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle the response (log or process further if needed)
	if resp.StatusCode != http.StatusCreated {
		return err
	}

	return nil
}
