package model

type HealthStatus struct {
	Message     string
	ServiceName string
	InstanceID  string
}

type HealthResponse struct {
	Message     string `json:"message" example:"ok"`
	ServiceName string `json:"service_name" example:"golang-backend-ebvn-k04-day-1"`
	InstanceID  string `json:"instance_id" example:"1234567890"`
}
