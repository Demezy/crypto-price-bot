package models

import (
	"fmt"
	"simple-go-telegram-bot/internal/consts"
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey"`
	TelegramId   string `gorm:"unique"`
	QueriyNumber int
	LastQuery    time.Time
	FirstQuery   time.Time
}

func (u User) String() string {
	return fmt.Sprintf("{id: %d, telegram_id: %s, query_number: %d, last_query: %s, first_query: %s}",
		u.ID,
		u.TelegramId,
		u.QueriyNumber,
		u.LastQuery.Format(consts.DateAndTime),
		u.FirstQuery.Format(consts.DateAndTime),
	)
}
