package bitcointrade_test

import (
	"bitcointrade"
	"testing"
)

func TestBalance(t *testing.T) {
	walletService := &bitcointrade.Wallet{}
	resp, err := walletService.Balance()
	if err != nil {
		t.Errorf("TestBalance => '%s'", err)
		return
	}

	if resp.Search("data").Data() == nil {
		t.Error("TestBalance => 'data is null'")
	}
}
