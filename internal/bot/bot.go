package bot

import (
	"log"
	"simple-go-telegram-bot/internal/consts"
	"simple-go-telegram-bot/internal/handlers"
	"simple-go-telegram-bot/internal/types"

	"gopkg.in/telebot.v3"
)

func StartBot(token string, db types.DB) {
	// configure bot
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: consts.BotPollingTimeout},
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
