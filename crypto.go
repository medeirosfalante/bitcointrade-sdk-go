package bitcointrade

import (
	"encoding/json"
	"fmt"
	"log"
)

// Crypto is Service for crypto manager
type Crypto struct {
	Crypto string
}

// WidthDraw is a struct for widthdral
type WidthDraw struct {
	Destination string  `json:"destination"`
	Amount      float64 `json:"amount"`
	Fee         float64 `json:"fee"`
	FeeType     string  `json:"fee_type"`
}

// Fee is func get estimative fee
func (c Crypto) Fee() (resp *Response, err error) {
	webClient := &WebClient{}
	resp, err = webClient.GET(fmt.Sprintf("%s/withdraw/fee", c.Crypto), nil)
	return
}

// ListWidthDraw is func get WidthDraw list
func (c Crypto) ListWidthDraw() (resp *Response, err error) {
	webClient := &WebClient{}
	resp, err = webClient.GET(fmt.Sprintf("%s/withdraw/fee", c.Crypto), nil)
	return
}

// ListDeposits is func get WidthDraw list
func (c Crypto) ListDeposits() (resp *Response, err error) {
	webClient := &WebClient{}
	resp, err = webClient.GET(fmt.Sprintf("%s/deposits", c.Crypto), nil)
	return
}

//CreateWidthDraw is func create a widthDraw list
func (c Crypto) CreateWidthDraw(widthDraw *WidthDraw) (resp *Response, erro error) {
	webClient := &WebClient{}
	widthDrawBytes, err := json.Marshal(widthDraw)
	if err != nil {
		log.Printf("err Marshal widthDraw: %s", err)
		return nil, err
	}
	resp, err = webClient.POST(fmt.Sprintf("%s/withdraw", c.Crypto), widthDrawBytes)
	return
}
