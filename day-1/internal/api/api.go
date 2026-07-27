package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/phuloicse/golang-backend-ebvn-k04/day-1/docs"
	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/config"
	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/handler"
	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(cfg *config.Config) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthService := service.NewHealthService(cfg.ServiceName, cfg.InstanceID)
	healthHandler := handler.NewHealthHandler(healthService)

	router.GET("/health-check", healthHandler.HealthCheck)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
