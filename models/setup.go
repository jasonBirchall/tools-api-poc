package models

import (
	"errors"

	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
)

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

func ConnectCache() {
	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	RedisClient = r
}
