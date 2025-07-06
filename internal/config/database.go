package config

import (
	"fmt"
	"github.com/nayeem-bd/Todo-App/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func ConnectDatabase(dbConfig DatabaseConfig) (*gorm.DB, error) {
	dns := buildDSN(dbConfig)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.MaxConnectionLifetime) * time.Second)

	logger.Info("Connected to PostgreSQL database")
	return db, nil
}

func buildDSN(dbConfig DatabaseConfig) string {
	options := ""
	for key, values := range dbConfig.Options {
		for _, value := range values {
			options += fmt.Sprintf("%s=%s ", key, value)
		}
	}
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d %s",
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
		options,
	)
}
