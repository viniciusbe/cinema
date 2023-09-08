package gormdb

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupGormDB() (*gorm.DB, error) {
	dsn := "host=localhost user=person password=person dbname=sunny port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("SetupGormDB -> %w", err)
	}

	return db, nil
}