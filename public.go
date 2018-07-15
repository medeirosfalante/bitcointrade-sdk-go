package bitcointrade

import (
	"fmt"
	"log"
)

// Public struct
type Public struct {
	Crypto string
}

// Ticker in exchange
func (p Public) Ticker() (*Response, error) {
	webClient := WebClient{}
	resp, err := webClient.GET(fmt.Sprintf("public/%s/ticker", p.Crypto), nil)
	if err != nil {
		log.Printf("error %s", err)
	}
	return resp, err
}

// Orders in exchange
func (p Public) Orders() (*Response, error) {
	webClient := WebClient{}
	resp, err := webClient.GET(fmt.Sprintf("public/%s/orders", p.Crypto), nil)
	if err != nil {
		log.Printf("error %s", err)
	}
	return resp, err
}

// Trades in exchange
func (p Public) Trades() (*Response, error) {
	webClient := WebClient{}
	resp, err := webClient.GET(fmt.Sprintf("public/%s/trades", p.Crypto), nil)
	if err != nil {
		log.Printf("error %s", err)
	}
	return resp, err
}
