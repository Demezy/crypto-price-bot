package handlers

import (
	"fmt"
	"simple-go-telegram-bot/internal/consts"
	"simple-go-telegram-bot/internal/types"
	"time"
)

func displayUserStat(u types.User) string {
	difference := time.Now().Sub(u.FirstQuery)

	return fmt.Sprintf(
		`query number: %d
last query: %s
first query: %s
days with us: %d`,
		u.QueriyNumber,
		u.LastQuery.Format(consts.DateAndTime),
		u.FirstQuery.Format(consts.DateAndTime),
		int(difference.Hours()/24),
	)
}
