package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/config"
	"github.com/mremperor-atwork/rpg1/api1/internal/database"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
	"github.com/mremperor-atwork/rpg1/api1/internal/routes"
	"github.com/mremperor-atwork/rpg1/api1/internal/supabase"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := supabase.Connect(cfg.SupabaseURL, cfg.SupabaseKey); err != nil {
		log.Fatalf("failed to connect to Supabase: %v", err)
	}
	defer supabase.Close()

	if err := database.Connect(cfg.DatabaseURL); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	routes.SetupAuthRoutes(r)
	routes.SetupSRDRoutes(r)
	routes.SetupCharacterRoutes(r)
	routes.SetupUserRoutes(r)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
