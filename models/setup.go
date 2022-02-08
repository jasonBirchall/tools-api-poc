package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	database, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Tools{})

	DB = database
}
