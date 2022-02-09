package models

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return err
	}

	err = database.AutoMigrate(&Tool{})
	if err != nil {
		return err
	}

	DB = database

	if DB == nil {
		return errors.New("No database connection")
	}

	return nil
}
