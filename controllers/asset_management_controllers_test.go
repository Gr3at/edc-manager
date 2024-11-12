package controllers

import "testing"

// import (
// 	"edc-proxy/services"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// // Mocking the service function
// func mockCreateResourcesSuccess(data map[string]interface{}) error {
// 	return nil // Simulate success
// }

// func mockCreateResourcesFailure(data map[string]interface{}) error {
// 	return errors.New("failed to create resources") // Simulate failure
// }

// func TestCreateOffering_Success(t *testing.T) {
// 	// Mock the CreateResources function to return success
// 	services.CreateResources = mockCreateResourcesSuccess

// 	// Set up the Gin router with the necessary route
// 	router := gin.Default()
// 	router.POST("/create-offering", CreateOffering)

// 	// Create a sample JSON request body
// 	requestBody := `{
// 		"resource": "example-resource-data"
// 	}`

// 	// Create an HTTP POST request
// 	req, _ := http.NewRequest("POST", "/create-offering", strings.NewReader(requestBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Use httptest to record the response
// 	w := httptest.NewRecorder()

// 	// Serve the HTTP request
// 	router.ServeHTTP(w, req)

// 	// Check if the status code is 200 OK
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	// Check if the response contains the expected success message
// 	assert.Contains(t, w.Body.String(), `"status":"Offering created"`)
// }

// func TestCreateOffering_Failure(t *testing.T) {
// 	// Mock the CreateResources function to return failure
// 	services.CreateResources = mockCreateResourcesFailure

// 	// Set up the Gin router with the necessary route
// 	router := gin.Default()
// 	router.POST("/create-offering", CreateOffering)

// 	// Create a sample JSON request body
// 	requestBody := `{
// 		"resource": "example-resource-data"
// 	}`

// 	// Create an HTTP POST request
// 	req, _ := http.NewRequest("POST", "/create-offering", strings.NewReader(requestBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Use httptest to record the response
// 	w := httptest.NewRecorder()

// 	// Serve the HTTP request
// 	router.ServeHTTP(w, req)

// 	// Check if the status code is 500 Internal Server Error
// 	assert.Equal(t, http.StatusInternalServerError, w.Code)

// 	// Check if the response contains the expected error message
// 	assert.Contains(t, w.Body.String(), `"error":"Failed to create offering"`)
// }

// func TestCreateOffering_InvalidJSON(t *testing.T) {
// 	// Set up the Gin router with the necessary route
// 	router := gin.Default()
// 	router.POST("/create-offering", CreateOffering)

// 	// Create an invalid JSON request body (missing closing brace)
// 	requestBody := `{
// 		"resource": "example-resource-data"`

// 	// Create an HTTP POST request
// 	req, _ := http.NewRequest("POST", "/create-offering", strings.NewReader(requestBody))
// 	req.Header.Set("Content-Type", "application/json")

// 	// Use httptest to record the response
// 	w := httptest.NewRecorder()

// 	// Serve the HTTP request
// 	router.ServeHTTP(w, req)

// 	// Check if the status code is 400 Bad Request
// 	assert.Equal(t, http.StatusBadRequest, w.Code)

// 	// Check if the response contains the expected error message
// 	assert.Contains(t, w.Body.String(), `"error":"EOF"`) // Or other error related to invalid JSON parsing
// }

func TestCreateAsset(t *testing.T) {

}
