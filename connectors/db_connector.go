package connectors

import (
	"fmt"
	"log"
	"lunar/conf"
	"lunar/types"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDatabase initializes the database connection.
func InitDatabase(cfg *conf.DatabaseConfig) *gorm.DB {
	once.Do(func() {
		var err error
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		log.Println("Connection established")

		// AutoMigrate models
		migrateErr := tableMigration(&types.RocketModel{})
		if migrateErr != nil {
			log.Fatalf("Failed to migrate tables: %v", migrateErr)
		}

	})
	return db
}

// Uses AutoMigrate to migrate models
func tableMigration(models ...interface{}) error {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	log.Println("Migrated successfully")
	return nil
}

// GetDB returns the database instance.
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database not initialized")
	}
	return db
}
