package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/mremperor-atwork/rpg1/api1/internal/database/seeds"
)

func main() {
	log.Println("🧪 Testing equipment items with nullable damage fields...")

	// Get database DSN from environment
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Connect to database without running migrations
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:        dsn,
		DriverName: "postgres",
	}), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Test weapon items (should have damage values)
	var weapons []seeds.EquipmentItem
	if err := db.Where("category = ?", "Weapon").Limit(5).Find(&weapons).Error; err != nil {
		log.Fatalf("Failed to query weapons: %v", err)
	}

	log.Println("📦 Weapon items (should have damage values):")
	for _, weapon := range weapons {
		damage := "nil"
		damageType := "nil"
		if weapon.Damage != nil {
			damage = fmt.Sprintf("%d", *weapon.Damage)
		}
		if weapon.DamageType != nil {
			damageType = *weapon.DamageType
		}
		log.Printf("   %s: Damage=%s, Type=%s", weapon.Name, damage, damageType)
	}

	// Test armor items (should have nil damage values)
	var armors []seeds.EquipmentItem
	if err := db.Where("category = ?", "Armor").Limit(5).Find(&armors).Error; err != nil {
		log.Fatalf("Failed to query armors: %v", err)
	}

	log.Println("🛡️  Armor items (should have nil damage values):")
	for _, armor := range armors {
		damage := "nil"
		damageType := "nil"
		if armor.Damage != nil {
			damage = fmt.Sprintf("%d", *armor.Damage)
		}
		if armor.DamageType != nil {
			damageType = *armor.DamageType
		}
		log.Printf("   %s: Damage=%s, Type=%s", armor.Name, damage, damageType)
	}

	// Test ammunition items (should have damage values)
	var ammo []seeds.EquipmentItem
	if err := db.Where("category = ?", "Ammunition").Limit(5).Find(&ammo).Error; err != nil {
		log.Fatalf("Failed to query ammunition: %v", err)
	}

	log.Println("🎯 Ammunition items (should have damage values):")
	for _, item := range ammo {
		damage := "nil"
		damageType := "nil"
		if item.Damage != nil {
			damage = fmt.Sprintf("%d", *item.Damage)
		}
		if item.DamageType != nil {
			damageType = *item.DamageType
		}
		log.Printf("   %s: Damage=%s, Type=%s", item.Name, damage, damageType)
	}

	log.Println("✅ Equipment items test completed!")
}
