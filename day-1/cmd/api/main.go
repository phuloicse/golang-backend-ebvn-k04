package main

import (
	"log"

	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/api"
	"github.com/phuloicse/golang-backend-ebvn-k04/day-1/internal/config"
)

// @title Golang Backend Health
// @version 1.0.0
// @description Health check API for Golang Backend EBVN K04 Day 1 assignment
// @BasePath /
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("faile to lead config: %v ", err)
	}

	router := api.SetupRouter(cfg)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("faile to start server: %v ", err)
	}

}
