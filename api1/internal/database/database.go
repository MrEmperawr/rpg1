package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(dsn string) error {
	var err error

	// Add PgBouncer compatibility parameters to the DSN
	// These parameters disable prepared statements which are not supported by PgBouncer
	if dsn != "" {
		log.Printf("Original DSN: %s", dsn)

		// Check if the DSN already has query parameters
		if len(dsn) > 0 && dsn[len(dsn)-1] != '?' {
			dsn += "?"
		} else if len(dsn) > 0 && dsn[len(dsn)-1] == '?' {
			// DSN already has query parameters, add & instead
			dsn = dsn[:len(dsn)-1] + "&"
		}

		// Add PgBouncer compatibility parameters
		dsn += "pgbouncer=true&prepared_statements=false&statement_cache_mode=describe"
		log.Printf("Modified DSN: %s", dsn)
	}

	gormConfig := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info), // Set to logger.Silent in production
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              false, // Disable prepared statements to avoid Supabase issues
		SkipDefaultTransaction:                   true,  // Skip default transaction to avoid connection issues
	}

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		DriverName:           "postgres", // Use lib/pq driver instead of pgx for better PgBouncer compatibility
		PreferSimpleProtocol: true,       // Use simple protocol to avoid prepared statement issues
	}), gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Configure connection pool for PgBouncer compatibility
	sqlDB.SetMaxOpenConns(1)    // Limit to 1 connection to avoid PgBouncer issues
	sqlDB.SetMaxIdleConns(1)    // Keep only 1 idle connection
	sqlDB.SetConnMaxLifetime(0) // Don't close connections automatically

	if err := sqlDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to database with GORM!")

	return nil
}

func GetDB() *gorm.DB {
	return DB
}

// RunMigrationsIfNeeded runs migrations only if explicitly requested
func RunMigrationsIfNeeded() error {
	if DB == nil {
		return fmt.Errorf("database not connected")
	}
	return RunMigrations(DB)
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
