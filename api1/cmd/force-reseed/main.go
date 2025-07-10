package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mremperor-atwork/rpg1/api1/internal/database"
	"github.com/mremperor-atwork/rpg1/api1/internal/database/seeds"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get database connection string
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Connect to database
	if err := database.Connect(dsn); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	db := database.GetDB()

	// Clear equipment items table
	log.Println("Clearing equipment_items table...")
	if err := db.Exec("DELETE FROM equipment_items").Error; err != nil {
		log.Fatalf("Failed to clear equipment_items table: %v", err)
	}

	// Re-seed equipment items
	log.Println("Re-seeding equipment items...")
	if err := seeds.SeedEquipmentItems(db); err != nil {
		log.Fatalf("Failed to seed equipment items: %v", err)
	}

	log.Println("Equipment items re-seeded successfully!")

	// Count the new items
	var count int64
	if err := db.Model(&seeds.EquipmentItem{}).Count(&count).Error; err != nil {
		log.Fatalf("Failed to count equipment items: %v", err)
	}

	log.Printf("Total equipment items in database: %d", count)

	// Count by category
	var categories []struct {
		Category string
		Count    int64
	}
	if err := db.Model(&seeds.EquipmentItem{}).Select("category, count(*) as count").Group("category").Scan(&categories).Error; err != nil {
		log.Fatalf("Failed to count by category: %v", err)
	}

	log.Println("\nEquipment items by category:")
	for _, cat := range categories {
		log.Printf("  %s: %d", cat.Category, cat.Count)
	}

	// Count by subcategory
	var subcategories []struct {
		Subcategory string
		Count       int64
	}
	if err := db.Model(&seeds.EquipmentItem{}).Select("subcategory, count(*) as count").Group("subcategory").Scan(&subcategories).Error; err != nil {
		log.Fatalf("Failed to count by subcategory: %v", err)
	}

	log.Println("\nEquipment items by subcategory:")
	for _, sub := range subcategories {
		log.Printf("  %s: %d", sub.Subcategory, sub.Count)
	}
}
