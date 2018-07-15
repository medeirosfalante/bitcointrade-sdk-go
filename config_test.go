package bitcointrade_test

import (
	"paymentbit/services/exchange/pkg/bitcointrade"
	"testing"
)

var (
	uriV1 = "https://api.bitcointrade.com.br/v1"
)

func TestGetPath(t *testing.T) {
	baseUri := bitcointrade.GetBasePath()
	if baseUri != uriV1 {
		t.Errorf("UrlBase not valid %s", baseUri)
	}

}
