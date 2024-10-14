package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateResources(t *testing.T) {
	// Mock a third-party service using httptest
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
	}))
	defer mockServer.Close()

	// Replace third-party service URL with mock server URL
	// originalURL := "https://third-party-service.com/api/resource"
	// mockURL := mockServer.URL
	// originalURL := mockURL // Simulate changing the URL for test

	data := map[string]interface{}{
		"resource": "some-resource-data",
	}

	err := CreateResources(data)
	assert.Nil(t, err, "Expected no error during resource creation")
}
