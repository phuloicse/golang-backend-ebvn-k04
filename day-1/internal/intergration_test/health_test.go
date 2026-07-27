package intergration_test

import (
	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/api"
	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthApiEndpoint(t *testing.T) {
	cfg := &config.Config{
		Port:        "8080",
		ServiceName: "golang-backend-ebvn-k04-day-1",
		InstanceID:  "1234567890",
	}

	router := api.SetupRouter(cfg)

	req := httptest.NewRequest("GET", "/health-check", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Code, http.StatusOK)
	assert.JSONEq(t, resp.Body.String(), `{
		"message":"ok",
		"service_name":"golang-backend-ebvn-k04-day-1",
		"instance_id":"1234567890"
		}`)
}
