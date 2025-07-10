package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mremperor-atwork/rpg1/api1/internal/database"
)

func main() {
	log.Println("🚀 FRESH START SCRIPT")
	log.Println("This will create a completely fresh database!")
	log.Println("")

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	// Get database connection string
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	log.Println("🔗 Connecting to database with lib/pq...")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("✅ Database connection successful")

	// Get current database name
	var dbName string
	err = db.QueryRow("SELECT current_database()").Scan(&dbName)
	if err != nil {
		log.Fatalf("Failed to get database name: %v", err)
	}
	log.Printf("📊 Current database: %s", dbName)

	// Get all tables in the current schema
	log.Println("🔍 Finding all tables to drop...")
	rows, err := db.Query(`
		SELECT tablename 
		FROM pg_tables 
		WHERE schemaname = 'public' 
		ORDER BY tablename
	`)
	if err != nil {
		log.Fatalf("Failed to query tables: %v", err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Failed to scan table name: %v", err)
		}
		tables = append(tables, tableName)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating tables: %v", err)
	}

	if len(tables) == 0 {
		log.Println("ℹ️  No tables found to drop")
	} else {
		log.Printf("📋 Found %d tables to drop:", len(tables))
		for _, table := range tables {
			log.Printf("   - %s", table)
		}

		log.Println("")
		log.Println("💥 Dropping all tables...")

		// Drop all tables with CASCADE to handle dependencies
		for _, table := range tables {
			log.Printf("   Dropping table: %s", table)
			_, err := db.Exec(fmt.Sprintf(`DROP TABLE IF EXISTS "%s" CASCADE`, table))
			if err != nil {
				log.Printf("⚠️  Warning: Failed to drop table %s: %v", table, err)
			} else {
				log.Printf("   ✅ Dropped table: %s", table)
			}
		}
	}

	log.Println("")
	log.Println("🧹 Database nuked successfully!")
	log.Println("")
	log.Println("🏗️  Running fresh migrations with GORM...")

	// Close the sql.DB connection and use GORM for migrations
	db.Close()

	// Use GORM to run migrations with a fresh connection
	gormDB, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect with GORM: %v", err)
	}

	// Run migrations
	if err := database.RunMigrations(gormDB); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("✅ Migrations completed successfully!")
	log.Println("")
	log.Println("🌱 Running fresh seeding...")

	// Run seeding
	if err := database.SeedDatabase(gormDB); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	log.Println("")
	log.Println("🎉 FRESH START COMPLETE!")
	log.Println("")
	log.Println("✨ Your database is now fresh and ready to go!")
}
