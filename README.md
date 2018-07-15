# bitcointrade-sdk-go

simple use example:

os.Setenv("USER_API_TOKEN", userToken)

marketService := &bitcointrade.Market{}

marketParams := &bitcointrade.MarketParams{
  Currency: "BTC",
}

resp, err := marketService.Summary(marketParams)

if err != nil {
  log.Printf("Summary => '%s'", err)
}

// Response based on https://github.com/Jeffail/gabs
resp.Search("data")
