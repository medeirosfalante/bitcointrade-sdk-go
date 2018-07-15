package bitcointrade

import (
	"log"
)

// Wallet struct
type Wallet struct {
}

// Balance is a func list balances
func (w Wallet) Balance() (*Response, error) {
	webClient := WebClient{}
	resp, err := webClient.GET("wallets/balance", nil)
	if err != nil {
		log.Printf("error %s", err)
	}
	return resp, err
}
