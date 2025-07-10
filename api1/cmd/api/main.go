package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/config"
	"github.com/mremperor-atwork/rpg1/api1/internal/database"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize database with GORM
	if err := database.Connect(cfg.DatabaseURL); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	r.Run(cfg.ServerAddress)
}
