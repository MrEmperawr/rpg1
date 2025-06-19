package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/config"
	"github.com/mremperor-atwork/rpg1/api1/internal/database"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize database with GORM
	if err := database.Connect(cfg.PostgresDSN); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	// Auto-migrate the database schema
	db := database.GetDB()
	if err := db.AutoMigrate(&models.User{}, &models.Character{}, &models.Item{}, &models.CharacterItem{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	log.Println("Database schema migrated successfully!")

	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	r.Run(cfg.ServerAddress)
}
