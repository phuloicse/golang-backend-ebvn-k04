package handler

import (
	"context"
	"net/http/httptest"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/model"
	"github.com/stretchr/testify/assert"
)

type mockHealthService struct{}

func (m *mockHealthService) GetHealthStatus(ctx context.Context) model.HealthStatus {
	return model.HealthStatus{
		Message:     "ok",
		ServiceName: "golang-backend-ebvn-k04-day-1",
		InstanceID:  "1234567890",
	}
}

func TestHealthHandler_HealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	healthService := &mockHealthService{}
	healthHandler := NewHealthHandler(healthService)

	router := gin.New()
	router.GET("/health-check", healthHandler.HealthCheck)

	req := httptest.NewRequest("GET", "/health-check", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, recorder.Code, http.StatusOK)
	assert.JSONEq(t, recorder.Body.String(), `{"message":"ok","service_name":"golang-backend-ebvn-k04-day-1","instance_id":"1234567890"}`)

}
