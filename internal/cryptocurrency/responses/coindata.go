package responses

type CoinData struct {
	MarketData struct {
		CurrentPrice struct {
			Usd float64 `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
}
