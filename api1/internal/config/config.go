package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SupabaseURL   string
	SupabaseKey   string
	ServerAddress string
	DatabaseURL   string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		SupabaseURL:   os.Getenv("SUPABASE_URL"),
		SupabaseKey:   os.Getenv("SUPABASE_KEY"),
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
	}
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
