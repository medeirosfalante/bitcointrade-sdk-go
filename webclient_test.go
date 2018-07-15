package bitcointrade_test

import (
	"bitcointrade"
	"encoding/json"
	"os"
	"testing"
)

const (
	userToken = "token"
)

func TestGET(t *testing.T) {
	webclient := bitcointrade.WebClient{}
	_, err := webclient.GET("public/BTC/ticker", nil)
	if err != nil {
		t.Errorf("erro test GET: %s", err)
	}

}

func TestPOST(t *testing.T) {
	webclient := bitcointrade.WebClient{}
	os.Setenv("USER_API_TOKEN", userToken)
	withdraw := &bitcointrade.WidthDraw{
		Destination: "mj9NZ98y1gwmgvxZytvqbMbUSo9M3E7CYN",
		Amount:      0.1,
		Fee:         0.11,
		FeeType:     "slow",
	}
	withdrawByte, _ := json.Marshal(&withdraw)
	_, err := webclient.POST("bitcoin/withdraw", withdrawByte)
	if err != nil {
		t.Errorf("erro test POST: %s", err)
	}

}
