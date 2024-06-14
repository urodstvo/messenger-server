package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgres() (*gorm.DB, error) {
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		return nil, fmt.Errorf("postgres error: no environment variable for POSTGRES_HOST")
	}

	port, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		return nil, fmt.Errorf("postgres error: no environment variable for POSTGRES_PORT")
	}

	user, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		return nil, fmt.Errorf("postgres error: no environment variable for POSTGRES_USER")
	}

	password, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		return nil, fmt.Errorf("postgres error: no environment variable for POSTGRES_PASSWORD")
	}

	name, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		return nil, fmt.Errorf("postgres error: no environment variable for POSTGRES_DB")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, name, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
