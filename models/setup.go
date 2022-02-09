package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dbURL := "postgres://pg:pass@localhost:5432/crud"

	// database, err := gorm.Open("sqlite3", "test.db")

	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&Tool{})
	if err != nil {
		panic(err)
	}

	DB = database
}
