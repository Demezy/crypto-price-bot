package cryptocurrency

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"simple-go-telegram-bot/internal/consts"
	"simple-go-telegram-bot/internal/cryptocurrency/responses"
)

func requestParse(url string, data any) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("api request error", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return errors.New("Currency not found")
	}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Println("api response parsing error", err)
		return err
	}
	return nil
}

func GetCurrencyPrice(currency string) (price float64, err error) {
	var data responses.CoinData
	err = requestParse(
		fmt.Sprintf(consts.CurrencyPrice, currency),
		&data,
	)
	return data.MarketData.CurrentPrice.Usd, err
}
