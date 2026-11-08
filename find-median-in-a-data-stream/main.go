package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"os"
)

func main() {
	{
		medianFinder := Constructor()
		medianFinder.AddNum(1)
		expect("#1 > #1", medianFinder.FindMedian(), 1)
		medianFinder.AddNum(3)
		expect("#1 > #2", medianFinder.FindMedian(), 2)
		medianFinder.AddNum(2)
		expect("#1 > #3", medianFinder.FindMedian(), 2)
	}
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

type MinHeap[T cmp.Ordered] []T

func (h MinHeap[T]) Len() int {
	return len(h)
}

func (h MinHeap[T]) Less(i, j int) bool {
	return cmp.Less(h[i], h[j])
}

func (h MinHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MinHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

type MaxHeap[T cmp.Ordered] []T

func (h MaxHeap[T]) Len() int {
	return len(h)
}

func (h MaxHeap[T]) Less(i, j int) bool {
	return cmp.Less(h[j], h[i])
}

func (h MaxHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap[T]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *MaxHeap[T]) Push(x any) {
	*h = append(*h, x.(T))
}

type MedianFinder struct {
	Small *MaxHeap[int]
	Large *MinHeap[int]
}

func Constructor() MedianFinder {
	medianFinder := MedianFinder{Small: &MaxHeap[int]{}, Large: &MinHeap[int]{}}
	heap.Init(medianFinder.Small)
	heap.Init(medianFinder.Large)
	return medianFinder
}

func (this *MedianFinder) AddNum(num int) {
	if this.Large.Len() > 0 {
		if num > (*this.Large)[0] {
			heap.Push(this.Large, num)
		} else {
			heap.Push(this.Small, num)
		}
	} else {
		heap.Push(this.Small, num)
	}
	if this.Small.Len() > this.Large.Len()+1 {
		heap.Push(this.Large, heap.Pop(this.Small))
	}
	if this.Large.Len() > this.Small.Len()+1 {
		heap.Push(this.Small, heap.Pop(this.Large))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.Small.Len() > this.Large.Len() {
		return float64((*this.Small)[0])
	}
	if this.Large.Len() > this.Small.Len() {
		return float64((*this.Large)[0])
	}
	return float64((*this.Small)[0]+(*this.Large)[0]) / 2.
}
