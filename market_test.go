package bitcointrade_test

import (
	"bitcointrade"
	"os"
	"testing"
)

func TestGetBookOrders(t *testing.T) {
	os.Setenv("USER_API_TOKEN", userToken)
	marketService := &bitcointrade.Market{}
	marketParams := &bitcointrade.MarketParams{
		Currency: "BTC",
	}
	resp, err := marketService.GetBookOrders(marketParams)
	if err != nil {
		t.Errorf("getBookOrders => '%s'", err)
		return
	}

	if resp.Search("data") == nil {
		t.Error("getBookOrders => 'data is null erro'")
	}
}
func TestSummary(t *testing.T) {
	os.Setenv("USER_API_TOKEN", userToken)
	marketService := &bitcointrade.Market{}
	marketParams := &bitcointrade.MarketParams{
		Currency: "BTC",
	}
	resp, err := marketService.Summary(marketParams)
	if err != nil {
		t.Errorf("Summary => '%s'", err)
		return
	}

	if resp.Search("data") == nil {
		t.Error("Summary => 'data is null erro'")
	}
}
func TestCreateOrder(t *testing.T) {
	os.Setenv("USER_API_TOKEN", userToken)
	marketService := &bitcointrade.Market{}
	order := &bitcointrade.Order{
		Currency:  "BTC",
		Amount:    0.5,
		Type:      "buy",
		SubType:   "market",
		UnitPrice: 1000,
	}

	resp, err := marketService.CreateOrder(order)
	if err != nil {
		t.Errorf("TestCreateOrder => '%s'", err)
		return
	}

	if resp.Search("message") == nil {
		t.Error("TestCreateOrder => 'message is null erro'")
	}

}
