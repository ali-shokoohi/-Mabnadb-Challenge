package controllers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shokoohi/mabnadp-challenge/models"
	"gitlab.com/shokoohi/mabnadp-challenge/routers"
)

// TestReturnAllInstruments - Unit test for ReturnAllInstruments controller
func TestReturnAllInstruments(t *testing.T) {
	// Get router object from our router package
	router := routers.GetRouter()
	// Create a http request
	request, err := http.NewRequest("GET", "/instruments", nil)
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 200, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))

	// Checks response type via convert response body type from []bytes to the []Instrument model type
	var result []models.Instrument
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Can't convert response body type to []Instrument model type")
}

// TestReturnSingleInstrument - Unit test fro ReturnSingleInstrument controller
func TestReturnSingleInstrument(t *testing.T) {
	// Get router object from our router package
	router := routers.GetRouter()
	// Create a http request
	request, err := http.NewRequest("GET", "/instruments/1", nil)
	// Check we have no error
	assert.Nil(t, err, "Request error must be Nil!")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	// Check response status code must be 200
	assert.Equal(t, 200, response.Result().StatusCode, fmt.Sprintf("Response code be: %d ! But It's: %d!!!", 200, response.Result().StatusCode))

	// Checks response type via convert response body type from []bytes to the []Instrument model type
	var result models.Instrument
	err = json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err, "Can't convert response body type to Instrument model type")
}
