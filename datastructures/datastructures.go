package datastructures

import (
	"container/list"
)

type Queue[T any] struct {
	queue list.List
}

// Add an Order to the queue
func (q *Queue[T]) Enqueue(order *T) {
	q.queue.PushBack(order)
}

// Remove and return the first Order
func (q *Queue[T]) Dequeue() *T {
	front := q.queue.Front()
	if front != nil {
		q.queue.Remove(front)
		if order, ok := front.Value.(*T); ok {
			return order
		}
	}
	return nil
}

// Peek at the first Order without removing
func (q *Queue[T]) Peek() *T {
	front := q.queue.Front()
	if front != nil {
		if order, ok := front.Value.(*T); ok {
			return order
		}
	}
	return nil
}

type Item struct {
	price float64
	queue Queue[any]
	index int
}

type MinHeap []*Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].price < h[j].price
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *MinHeap) Push(x any) {
	n := len(*h)
	item := x.(*Item)
	item.index = n
	*h = append(*h, item)
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*h = old[0 : n-1]
	return item
}

// TODO: Check if this will be needed
// func (h *MinHeap) Update()

type MaxHeap []*Item

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].price > h[j].price
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *MaxHeap) Push(x any) {
	n := len(*h)
	item := x.(*Item)
	item.index = n
	*h = append(*h, item)
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*h = old[0 : n-1]
	return item
}
