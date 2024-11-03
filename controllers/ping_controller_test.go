package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router for testing
	router := gin.Default()
	router.GET("/ping", Ping)

	// Create a new HTTP request for the /ping route
	req, _ := http.NewRequest("GET", "/ping", nil)
	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request using the router
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusOK, w.Code)

	// Assert the response body
	expectedBody := `{"message":"pong"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}
