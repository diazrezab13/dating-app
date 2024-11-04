package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var token string
var id float64

type Input struct {
	ProfileID uint `json:"profile_id"`
	Like      bool `json:"like"`
}

type InputUpgrade struct {
	UserID uint `json:"user_id"`
}

func TestSignUpProfile(t *testing.T) {
	router := setupTestRouter()
	setupUser()

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

func TestLoginProfile(t *testing.T) {
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

	token = responseBody["token"].(string)
}

func TestGetProfile(t *testing.T) {
	router := setupTestRouter()

	// Create a new HTTP POST request
	req, _ := http.NewRequest("GET", "/profile", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body
	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	users := responseBody["profiles"].([]interface{})
	data := users[0].(map[string]interface{})
	id = data["ID"].(float64)
}

func TestSwipeSuccess(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := Input{
		ProfileID: uint(id),
		Like:      true,
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/profile/swipe", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the response body
	var responseBody map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
}

func TestSwipeAlreadySwiped(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := Input{
		ProfileID: uint(id),
		Like:      true,
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/profile/swipe", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder
	w := httptest.NewRecorder()
	for i := 1; i < 3; i++ {
		router.ServeHTTP(w, req)
	}

	// Assert the response code
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpgradePremiumFailed(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := InputUpgrade{
		UserID: uint(id + 1),
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/profile/upgrade", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestUpgradePremium(t *testing.T) {
	router := setupTestRouter()

	// Sample request body
	requestBody := InputUpgrade{
		UserID: uint(id),
	}
	jsonValue, _ := json.Marshal(requestBody)

	// Create a new HTTP POST request
	req, _ := http.NewRequest("POST", "/profile/upgrade", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Create a response recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the response code
	assert.Equal(t, http.StatusOK, w.Code)
}
