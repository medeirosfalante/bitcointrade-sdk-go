package bitcointrade_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rafaeltokyo/bitcointrade-sdk-go"
)

func TestTicket(t *testing.T) {
	godotenv.Load()
	client := bitcointrade.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Public().Ticker("BRLBTC")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}

func TestOrders(t *testing.T) {
	godotenv.Load()
	client := bitcointrade.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Public().Orders("BRLBTC")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}

func TestTrades(t *testing.T) {
	godotenv.Load()
	client := bitcointrade.New("", os.Getenv("ENV"))
	response, errAPI, err := client.Public().Trades("BRLBTC")
	if err != nil {
		t.Errorf("err : %s", err)
		return
	}
	if errAPI != nil {
		t.Errorf("errAPI : %#v", errAPI)
		return
	}
	if response == nil {
		t.Error("response is null")
		return
	}
}
