package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig_NewConfig(t *testing.T) {
	t.Setenv("APP_PORT", "8080")
	t.Setenv("SERVICE_NAME", "golang-backend-ebvn-k04-day-1")
	t.Setenv("INSTANCE_ID", "1234567890")

	cfg, err := NewConfig()

	require.NoError(t, err)
	require.NotNil(t, cfg)

	assert.Equal(t, cfg.Port, "8080")
	assert.Equal(t, cfg.ServiceName, "golang-backend-ebvn-k04-day-1")
	assert.Equal(t, cfg.InstanceID, "1234567890")
}
