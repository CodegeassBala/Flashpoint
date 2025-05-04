package datastructures

import (
	"container/heap"
	"slices"

	"testing"
)

func TestQueue(t *testing.T) {
	queue := Queue[int]{}
	// Peek onto an empty queue should return nil
	t1 := queue.Peek()
	if t1 != nil {
		t.Errorf("Peek on empty queue did not return nil")
	}
	// enqueue
	v1 := 24
	queue.Enqueue(&v1)
	v2 := 25
	queue.Enqueue(&v2)
	if *queue.Peek() != v1 {
		t.Errorf("Peek did not return the elements in order")
	}
	r1 := queue.Dequeue()
	if *r1 != v1 {
		t.Errorf("Dequeue did not return the elements in order")
	}
}

func TestMinHeap(t *testing.T) {
	miniHeap := MinHeap{}
	prices := []float64{34.34, 189.12, 2.0, 27.013, 43.444, 189.05, 34.12, 2.5, 27.67}
	sorted_prices := make([]float64, len(prices))
	copy(sorted_prices, prices)
	slices.Sort(sorted_prices)
	for i, p := range prices {
		heap.Push(&miniHeap, &Item{
			price: p,
			queue: Queue[any]{},
			index: i,
		})
	}
	poped_prices := []float64{}
	for miniHeap.Len() > 0 {
		poped_item := heap.Pop(&miniHeap).(*Item)
		poped_prices = append(poped_prices, float64(poped_item.price))
	}
	for i := range sorted_prices {
		if sorted_prices[i] != poped_prices[i] {
			t.Errorf("MinHeap 's pop order does not match")
		}
	}
}

func TestMaxHeap(t *testing.T) {
	maxiHeap := MaxHeap{}
	prices := []float64{34.34, 189.12, 2.0, 27.013, 43.444, 189.05, 34.12, 2.5, 27.67}
	sorted_prices := make([]float64, len(prices))
	copy(sorted_prices, prices)
	slices.Sort(sorted_prices)
	slices.Reverse(sorted_prices)
	for i, p := range prices {
		heap.Push(&maxiHeap, &Item{
			price: p,
			queue: Queue[any]{},
			index: i,
		})
	}
	poped_prices := []float64{}
	for maxiHeap.Len() > 0 {
		poped_item := heap.Pop(&maxiHeap).(*Item)
		poped_prices = append(poped_prices, float64(poped_item.price))
	}
	for i := range sorted_prices {
		if sorted_prices[i] != poped_prices[i] {
			t.Errorf("MaxHeap 's pop order does not match")
		}
	}
}
