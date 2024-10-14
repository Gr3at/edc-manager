package routes

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoutesIntegration(t *testing.T) {
	router := gin.Default()

	// Set up routes
	SetupRoutes(router)

	// Test the GET /connectors endpoint
	t.Run("GET /connectors", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/connectors", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	// Test the POST /connectors endpoint
	t.Run("POST /create-record", func(t *testing.T) {
		payload := `{
			"api_url": "https://example.com/api",
			"api_key": "some-secret-key",
			"sub_id": "sub123",
			"available_to_all_org_users": true
		}`

		req, _ := http.NewRequest("POST", "/connectors", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), `"api_url":"https://example.com/api"`)
	})
}
