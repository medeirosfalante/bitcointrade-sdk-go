package bitcointrade

import (
	"encoding/json"
	"log"
)

// Order struct for a new order body
type Order struct {
	Currency  string  `json:"currency"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	SubType   string  `json:"subtype"`
	UnitPrice float64 `json:"unit_price"`
}

// Market struct operattion
type Market struct {
}

// MarketParams i a struct for getBook any and getBook user
type MarketParams struct {
	Currency    string `json:"currency,omitempty"`
	Status      string `json:"status,omitempty"`
	StartDate   string `json:"start_date,omitempty"`
	EndDate     string `json:"end_date,omitempty"`
	Type        string `json:"type,omitempty"`
	PageSize    int32  `json:"page_size,omitempty"`
	CurrentPage int32  `json:"current_page,omitempty"`
}

// GetBookOrders func
func (m Market) GetBookOrders(marketParams *MarketParams) (resp *Response, err error) {
	webClient := &WebClient{}
	marketParamsBytes, err := json.Marshal(marketParams)
	if err != nil {
		log.Printf("erro Marshel marketParams: %s", err)
	}

	resp, err = webClient.GET("market", marketParamsBytes)
	return

}

// Summary func
func (m Market) Summary(marketParams *MarketParams) (resp *Response, err error) {
	webClient := &WebClient{}

	marketParamsBytes, err := json.Marshal(marketParams)
	if err != nil {
		log.Printf("erro Marshel marketParams: %s", err)
	}

	resp, err = webClient.GET("market/summary", marketParamsBytes)
	return

}

// CreateOrder is a func create a new order in Book
func (m Market) CreateOrder(order *Order) (resp *Response, err error) {
	webClient := &WebClient{}
	orderBytes, err := json.Marshal(order)
	if err != nil {
		log.Printf("erro Marshel Order: %s", err)
	}
	resp, err = webClient.POST("market/create_order", orderBytes)
	return
}
