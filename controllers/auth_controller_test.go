package controllers

import (
	"bytes"
	"dating-app/config"
	"dating-app/middlewares"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var user, email, password string

// Load the .env file before running any tests
func TestMain(m *testing.M) {
	// Load the .env file
	if err := godotenv.Load("../.env"); err != nil {
		// Log a warning if the .env file is not found, but continue running the tests
		println("Warning: .env file not found")
	}

	setupUser()

	// Run the tests
	os.Exit(m.Run())
}

// Helper function to create a test router with our endpoints
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	config.ConnectDatabase()
	config.ConnectRedis()

	router.POST("/signup", SignUp)
	router.POST("/login", Login)
	router.Use(middlewares.AuthMiddleware())
	router.GET("/profile", GetProfiles)
	router.POST("/profile/swipe", Swipe)
	router.POST("/profile/upgrade", UpgradePremium)

	return router
}

// Helper function to generate a random string of a given length
func randomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// Helper function to generate a random email
func randomEmail() string {
	return fmt.Sprintf("%s@example.com", randomString(8))
}

// Helper function to generate a random password
func randomPassword() string {
	return randomString(12) // Generate a 12-character password
}

// Helper function to create a test router with our endpoints
func setupUser() {
	user = randomString(12)
	email = randomEmail()
	password = randomPassword()
}

func TestSignUpRequired(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := map[string]string{
		"username": user,
		"email":    email,
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSignUp(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := map[string]string{
		"username": user,
		"email":    email,
		"password": password,
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body
	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Contains(t, responseBody, "user")
}

func TestSignUpFailed(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := map[string]string{
		"username": user,
		"email":    email,
		"password": password,
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestLoginRequired(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := map[string]string{
		"email": email,
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body
	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Contains(t, responseBody, "token")
}

func TestLoginFailed(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := map[string]string{
		"email":    "john@example.com",
		"password": "john123",
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
