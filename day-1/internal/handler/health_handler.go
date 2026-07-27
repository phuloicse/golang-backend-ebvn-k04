package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/model"
	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/service"
)

type HealthHandler struct {
	healthService service.HealthService
}

func NewHealthHandler(healthService service.HealthService) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
}

// HealthCheck godoc
// @Summary Health check
// @Description Health check API for Golang Backend EBVN K04 Day 1 assignment
// @Tags health
// @Produce json
// @Success 200 {object} model.HealthResponse
// @Router /health-check [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	result := h.healthService.GetHealthStatus(c.Request.Context())

	c.JSON(http.StatusOK, model.HealthResponse{
		Message:     result.Message,
		ServiceName: result.ServiceName,
		InstanceID:  result.InstanceID,
	})
}
