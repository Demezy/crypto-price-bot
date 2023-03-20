package database

import (
	"errors"
	"fmt"
	"os"
	models "simple-go-telegram-bot/internal/db_models"
	"simple-go-telegram-bot/internal/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(host, user, password, port string) (db types.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		host,
		user,
		password,
		port,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateDB(db types.DB) error {
	return db.AutoMigrate(&models.User{})
}

func AutoDBSetup() (db types.DB, err error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	if host == "" || port == "" {
		return nil, errors.New("Undefined database host or port")
	}
	db, err = ConnectDB(host, user, password, port)
	if err != nil {
		return nil, err
	}
	return db, MigrateDB(db)
}
