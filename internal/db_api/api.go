package api

import (
	"errors"
	models "simple-go-telegram-bot/internal/db_models"
	"simple-go-telegram-bot/internal/types"
	"time"
)

type tID = string
type User = models.User

func GetUser(db types.DB, telegramId tID) (user User, err error) {
	err = db.Where("telegram_id = ?", telegramId).First(&user).Error
	return user, err
}

func CreateUser(db types.DB, telegramId tID) (user User, err error) {
	user.TelegramId = telegramId
	user.FirstQuery = time.Now()
	err = db.Create(&user).Error
	return user, err
}

func UpdateUser(db types.DB, user User) error {
	return db.Model(&user).Updates(user).Error
}

func TrackUserQuery(db types.DB, telegramId tID) error {
	user, err := GetUser(db, telegramId)
	if err != nil {
		var err2 error
		user, err2 = CreateUser(db, telegramId)
		if err2 != nil {
			return errors.Join(err, err2)
		}
	}
	user.LastQuery = time.Now()
	user.QueriyNumber++
	return UpdateUser(db, user)
}
