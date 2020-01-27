package main_test

import (
	"encoding/json"
	"hostgatorBackend/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var routerMain = router.StartRouter()

func TestRoutePricesAllSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/prices", nil)
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestRoutePricesOneSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/prices/5", nil)
	response := executeRequest(req)
	assert.Equal(t, http.StatusOK, response.Code)

	if assert.NotNil(t, response.Body) {
		var bodyRequest map[string]interface{}
		_ = json.NewDecoder(response.Body).Decode(&bodyRequest)
		assert.Equal(t, "Plano P", bodyRequest["name"])

	}
}

func TestRoutePricesOneNotExists(t *testing.T) {
	req, _ := http.NewRequest("GET", "/prices/1", nil)
	response := executeRequest(req)
	assert.Equal(t, http.StatusBadRequest, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	routerMain.ServeHTTP(rr, req)
	return rr
}
