package main

import (
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"

	"simple-go-telegram-bot/internal/db"
	"simple-go-telegram-bot/internal/handlers"
	"simple-go-telegram-bot/internal/types"
)

func StartBot(token string, db types.DB) {
	// configure bot
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 400 * time.Millisecond},
		OnError: func(err error, ctx telebot.Context) {
			log.Println(err)
		},
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	handlers.SetupHandlers(bot, db)
	bot.Start()
}

func main() {
	token := os.Getenv("TG_TOKEN")
  db, err := database.AutoDBSetup()
  if err!=nil{
    log.Fatal(err)
  }
	StartBot(token, db)
}