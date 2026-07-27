package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealthStatus_GetHealthStatus(t *testing.T) {
	healthService := NewHealthService("golang-backend-ebvn-k04-day-1", "1234567890")

	result := healthService.GetHealthStatus(context.Background())

	assert.Equal(t, result.Message, "ok")
	assert.Equal(t, result.ServiceName, "golang-backend-ebvn-k04-day-1")
	assert.Equal(t, result.InstanceID, "1234567890")
}
