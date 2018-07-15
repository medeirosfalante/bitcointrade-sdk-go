# bitcointrade-sdk-go

simple use example:

os.Setenv("USER_API_TOKEN", userToken)

marketService := &bitcointrade.Market{}

marketParams := &bitcointrade.MarketParams{
  Currency: "BTC",
}

resp, err := marketService.Summary(marketParams)

// Response based on https://github.com/Jeffail/gabs
