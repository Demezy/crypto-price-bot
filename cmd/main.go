package main

import (
	"log"
	"os"
	"simple-go-telegram-bot/internal/bot"
	"simple-go-telegram-bot/internal/consts"
	database "simple-go-telegram-bot/internal/db"
)

func main() {
	token := os.Getenv(consts.TgToken)
	db, err := database.AutoDBSetup()
	if err != nil {
		log.Fatal(err)
	}
	bot.StartBot(token, db)
}
