package bitcointrade_test

import (
	"bitcointrade"
	"testing"
)

func TestTicket(t *testing.T) {
	publicService := &bitcointrade.Public{
		Crypto: "BTC",
	}
	resp, err := publicService.Ticker()
	if err != nil {
		t.Errorf("TestTicket => '%s'", err)
		return
	}
	if resp.Search("data") == nil {
		t.Error("TestTicket => 'data is null erro'")
	}
}

func TestOrders(t *testing.T) {
	publicService := &bitcointrade.Public{
		Crypto: "BTC",
	}
	resp, err := publicService.Orders()
	if err != nil {
		t.Errorf("TestOrders => '%s'", err)
		return
	}
	if resp.Search("data") == nil {
		t.Error("TestOrders => 'data is null'")
	}
}
func TestTrades(t *testing.T) {
	publicService := &bitcointrade.Public{
		Crypto: "BTC",
	}
	resp, err := publicService.Trades()
	if err != nil {
		t.Errorf("TestTrades => '%s'", err)
		return
	}
	if resp.Search("data") == nil {
		t.Error("TestOrders => 'data is null'")
	}
}
