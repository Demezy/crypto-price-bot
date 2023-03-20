package handlers

import (
	"errors"
	"fmt"
	"log"
	"simple-go-telegram-bot/internal/consts"
	"simple-go-telegram-bot/internal/cryptocurrency"
	api "simple-go-telegram-bot/internal/db_api"
	"simple-go-telegram-bot/internal/types"

	"gopkg.in/telebot.v3"
)

func setStartHandler(bot *telebot.Bot) {
	bot.Handle(
		consts.CommandStart,
		func(ctx telebot.Context) error {
			ctx.Send("Welcome!")
			return nil
		})
}
func queryInfoHandler(bot *telebot.Bot, db types.DB) {
	bot.Handle(
		consts.CommandQuery,
		func(ctx telebot.Context) error {
			if len(ctx.Args()) != 1 {
				ctx.Send("Undefined number of arguments")
				return errors.New("Sent wrong number of arguments")
			}
			price, err := cryptocurrency.GetCurrencyPrice(ctx.Args()[0])
			if err != nil {
				ctx.Send("Provide valid cryptocurrency name")
				return errors.New("Invalid currency name")
			}
			defer api.TrackUserQuery(db, fmt.Sprint(ctx.Sender().ID))
			message := fmt.Sprintf(
				consts.CurrencyPriceFormat,
				ctx.Args()[0],
				price,
			)
			ctx.Send(message)
			return nil
		})
}

func setInfoHandler(bot *telebot.Bot, db types.DB) {
	bot.Handle(consts.CommandInfo,
		func(ctx telebot.Context) error {
			user, err := api.GetUser(db, fmt.Sprint(ctx.Sender().ID))
			if err != nil {
				ctx.Send("You have no statistics yet")
				return nil
			}
			ctx.Send(user.String())
			return nil
		})
}

func setEchoHandler(bot *telebot.Bot) {
	bot.Handle(
		telebot.OnText,
		func(ctx telebot.Context) error {
			log.Println(ctx.Sender().ID, ctx.Text())
			ctx.Send(ctx.Text())
			return nil
		},
	)
}

func SetupHandlers(bot *telebot.Bot, db types.DB) {
	setInfoHandler(bot, db)
	queryInfoHandler(bot, db)
	setStartHandler(bot)
	setEchoHandler(bot)
}
