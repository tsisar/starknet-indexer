package storage

import (
	"fmt"
	"github.com/tsisar/starknet-indexer/internal/storage/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Gorm wraps a GORM database instance.
type Gorm struct {
	DB *gorm.DB
}

// InitGorm establishes a connection to the PostgreSQL database using configuration values.
// It also ensures the required schemas (e.g., "core") are created.
func InitGorm() (*Gorm, error) {
	db, err := gorm.Open(postgres.Open(DSNFromConfig()), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get raw db failed: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("db ping failed: %w", err)
	}

	// Create the core schema if it doesn't exist
	for _, schema := range []string{"starknet"} {
		if err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schema)).Error; err != nil {
			return nil, fmt.Errorf("failed to create schema %s: %w", schema, err)
		}
	}

	if err := db.AutoMigrate(
		&model.Status{},
		&model.Block{},
		&model.Contract{},
		&model.Transaction{},
		&model.Event{},
	); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return &Gorm{DB: db}, nil
}

// Close closes the underlying SQL database connection.
func (g *Gorm) Close() error {
	sqlDB, err := g.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}
	return sqlDB.Close()
}
