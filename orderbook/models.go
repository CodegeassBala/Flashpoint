package orderbook

import (
	"container/list"
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

type Queue struct {
	queue list.List
}

// Add an Order to the queue
func (q *Queue) Enqueue(order *Order) {
	q.queue.PushBack(order)
}

// Remove and return the first Order
func (q *Queue) Dequeue() *Order {
	front := q.queue.Front()
	if front != nil {
		q.queue.Remove(front)
		if order, ok := front.Value.(*Order); ok {
			return order
		}
	}
	return nil
}

// Peek at the first Order without removing
func (q *Queue) Peek() *Order {
	front := q.queue.Front()
	if front != nil {
		if order, ok := front.Value.(*Order); ok {
			return order
		}
	}
	return nil
}

type PriceMap struct {
	Price float64 `json:"price"`
	Queue Queue   `json:"queue"`
}

type Ledger struct {
	PriceMap []PriceMap `json:"price_map"`
	Side     Side       `json:"side"` // "buy" or "sell"s
}

type OrderBook struct {
	BuyLedger  Ledger `json:"buy_ledger"`
	SellLedger Ledger `json:"sell_ledger"`
	Ticker     string `json:"ticker"`
}

func (ob *OrderBook) Initialize() {
	ob.BuyLedger = Ledger{
		PriceMap: []PriceMap{},
		Side:     Buy,
	}
	ob.SellLedger = Ledger{
		PriceMap: []PriceMap{},
		Side:     Sell,
	}
	return
}
