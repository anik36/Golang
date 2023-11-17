package token

import (
	"bytes"
	"encoding/json"
	"go-framework/internal/wscutils"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// Response represents the standard structure of a response of the web service.
type Response struct {
	Status   string                  `json:"status"`
	Data     any                     `json:"data"`
	Messages []wscutils.ErrorMessage `json:"messages"`
}

func TestMain(m *testing.M) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)
	// Run the tests
	os.Exit(m.Run())
}

func TestGetRealmAdminToken(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and register your routes
	r := gin.Default()
	RegisterTokenHandlers(r)

	// Create a request to send to the above route
	var jsonStr = `{
		"ver" : 1,
		"data": {
		"username": "user_1",
		"password": "user",
		"realm": "remiges-tech"
		}
	}`
	req, _ := http.NewRequest(http.MethodPost, "/login-admin", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	// Perform the request
	r.ServeHTTP(w, req)

	// Check the status code is what you expect
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected to get status %d but instead got %d. Body: %s\n", http.StatusOK, w.Code, w.Body.String())
	}
}

func TestGetRealmAdminTokenFailure(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)
	// Setup your router, just like you did in your main function, and register your routes
	r := gin.Default()
	RegisterTokenHandlers(r)
	// Create a request with invalid user data to send to the above route
	var jsonStr = []byte(`{
		"ver" : 1,
		"data": {
		"password": "user",
		"realm": ""
		}
	}`)
	req, _ := http.NewRequest(http.MethodPost, "/login-admin", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	// Perform the request
	r.ServeHTTP(w, req)
	// Check the status code is what you expect
	if w.Code != http.StatusOK {
		t.Errorf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
	// Examine the response and check the error messages
	// Deserialize response body
	var response Response
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Failed to decode response body: %s\n", err.Error())
		return
	}
	// Expect two validation errors (one for the email and one for the age)
	if len(response.Messages) != 2 {
		t.Errorf("Expected 2 validation errors but got %d\n", len(response.Messages))
	}
}
