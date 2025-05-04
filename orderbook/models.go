package orderbook

import (
	"flashpoint/datastructures"
)

type Side int

const (
	Buy Side = iota
	Sell
)

var sideToString = map[Side]string{
	Buy:  "buy",
	Sell: "sell",
}

func (s Side) String() string {
	return sideToString[s]
}

type Order struct {
	ID        *string `json:"id"`
	Price     float64 `json:"price"`
	Units     float64 `json:"units"`
	Side      Side    `json:"side"` // "buy" or "sell"
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}

type OrderBook struct {
	BuyLedger  datastructures.MaxHeap `json:"buy_ledger"`
	SellLedger datastructures.MinHeap `json:"sell_ledger"`
	Ticker     string                 `json:"ticker"`
}

func Initialize(ticker string) *OrderBook {
	return &OrderBook{
		Ticker:     ticker,
		BuyLedger:  datastructures.MaxHeap{},
		SellLedger: datastructures.MinHeap{},
	}
}
