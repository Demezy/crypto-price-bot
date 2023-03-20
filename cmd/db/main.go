package main

import (
	"fmt"
	"log"
	"simple-go-telegram-bot/internal/db"
	api "simple-go-telegram-bot/internal/db_api"
	"time"
)

func main() {
	db, err := database.AutoDBSetup()
	if err != nil {
		log.Fatal(err)
	}
	database.MigrateDB(db)
	const tid = "testid"

	api.CreateUser(db, tid)
	user, err := api.GetUser(db, tid)
	fmt.Println(user, err)
	time.Sleep(2 * time.Second)
	api.TrackUserQuery(db, tid)
	user, err = api.GetUser(db, tid)
	fmt.Println(user, err)
}
