package bitcointrade

import (
	"fmt"
)

// Public struct
type Public struct {
	client *APIClient
	Crypto string
}

//Ticker - return a ticket exchange
type Ticker struct {
	High     float32 `json:"high"`
	Low      float32 `json:"low"`
	Volume   float64 `json:"volume"`
	Quantity int32   `json:"quantity"`
	Last     float32 `json:"last"`
	Buy      float32 `json:"buy"`
	Sell     float32 `json:"sell"`
	Date     string  `json:"date"`
}

//Order - Struct for order
type Order struct {
	UnitPrice float32 `json:"unit_price"`
	Code      string  `json:"code"`
	Amount    float64 `json:"amount"`
}

//Orders - Struct for all list order
type Orders struct {
	Bids []Order `json:"bids"`
	Asks []Order `json:"asks"`
}

//TraderPagination - pagination in trade
type TraderPagination struct {
	CurrentPage    int32 `json:"current_page"`
	PageSize       int32 `json:"page_size"`
	RegistersCount int32 `json:"registers_count"`
	TotalPages     int32 `json:"total_pages"`
}

//Trader - struct for trade
type Trader struct {
	Type             string  `json:"type"`
	Amount           float64 `json:"amount"`
	UnitPrice        float32 `json:"unit_price"`
	ActiveOrderCode  string  `json:"active_order_code"`
	PassiveOrderCode string  `json:"passive_order_code"`
	Date             string  `json:"date"`
}

//Traders - Trader response
type Traders struct {
	Pagination TraderPagination `json:"pagination"`
	Trades     []Trader         `json:"trades"`
}

//TradeQuery - trade query
type TradeQuery struct {
	PageSize    int32 `json:"page_size"`
	CurrentPage int32 `json:"current_page"`
}

//Public - Create a new instance struct
func (c *APIClient) Public() *Public {
	return &Public{client: c}
}

// Ticker in exchange
func (p Public) Ticker(pair string) (*Ticker, *Error, error) {
	var response *Ticker
	err, errAPI := p.client.Request("GET", fmt.Sprintf("public/%s/ticker", pair), nil, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

// Orders - Orders in exchange
func (p Public) Orders(pair string) (*Orders, *Error, error) {
	var response *Orders
	err, errAPI := p.client.Request("GET", fmt.Sprintf("public/%s/orders", pair), nil, nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

// Trades - Trades in exchange
func (p Public) Trades(pair string) (*Traders, *Error, error) {
	var response *Traders
	err, errAPI := p.client.Request("GET", fmt.Sprintf("public/%s/trades", pair), nil, &TradeQuery{PageSize: 10, CurrentPage: 1}, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
