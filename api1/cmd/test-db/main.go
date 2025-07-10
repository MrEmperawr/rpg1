package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mremperor-atwork/rpg1/api1/internal/database"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	fmt.Println("Testing database connection...")
	fmt.Printf("Database URL: %s\n", dbURL)

	// Test connection
	if err := database.Connect(dbURL); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("✅ Database connection successful!")

	// Get database instance and test a simple query
	db := database.GetDB()

	// Test if we can query the existing tables
	var userCount int64
	if err := db.Table("user").Count(&userCount).Error; err != nil {
		fmt.Printf("⚠️  Could not query 'user' table: %v\n", err)
	} else {
		fmt.Printf("✅ 'user' table accessible, count: %d\n", userCount)
	}

	var characterCount int64
	if err := db.Table("character").Count(&characterCount).Error; err != nil {
		fmt.Printf("⚠️  Could not query 'character' table: %v\n", err)
	} else {
		fmt.Printf("✅ 'character' table accessible, count: %d\n", characterCount)
	}

	var speciesCount int64
	if err := db.Table("species").Count(&speciesCount).Error; err != nil {
		fmt.Printf("⚠️  Could not query 'species' table: %v\n", err)
	} else {
		fmt.Printf("✅ 'species' table accessible, count: %d\n", speciesCount)
	}

	fmt.Println("\n🎉 Database connection test completed!")

	// Close connection
	database.Close()
}
