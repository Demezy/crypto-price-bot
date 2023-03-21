package types

import (
	models "simple-go-telegram-bot/internal/db_models"

	"gorm.io/gorm"
)

type DB = *gorm.DB
type User = models.User
