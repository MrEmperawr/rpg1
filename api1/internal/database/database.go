package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(dsn string) error {
	var err error

	gormConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info), // Set to logger.Silent in production
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              false, // Disable prepared statements to avoid Supabase issues
	}

	// Configure postgres driver to use pgx
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:        dsn,
		DriverName: "pgx", // Use pgx driver
	}), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database with GORM!")

	if err := RunMigrations(DB); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func Close() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}
