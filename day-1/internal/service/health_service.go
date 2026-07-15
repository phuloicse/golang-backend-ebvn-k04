package service

import (
	"context"

	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/model"
)

type HealthService interface {
	GetHealthStatus(ctx context.Context) model.HealthStatus
}

type healthService struct {
	serviceName string
	instanceID  string
}

func NewHealthService(serviceName string, instanceID string) HealthService {
	return &healthService{
		serviceName: serviceName,
		instanceID:  instanceID,
	}
}

func (s *healthService) GetHealthStatus(ctx context.Context) model.HealthStatus {
	return model.HealthStatus{
		Message:     "ok",
		ServiceName: s.serviceName,
		InstanceID:  s.instanceID,
	}
}
