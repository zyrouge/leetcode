package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	expect("#1", minMeetingRooms([]Interval{{0, 40}, {5, 10}, {15, 20}}), 2)
	expect("#2", minMeetingRooms([]Interval{{4, 9}}), 1)
	expect("#2", minMeetingRooms([]Interval{{1, 5}, {2, 6}, {3, 7}, {4, 8}, {5, 9}}), 4)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

type Interval struct {
	start int
	end   int
}

// type MinHeap[T cmp.Ordered] []T

// func (h MinHeap[T]) Len() int {
// 	return len(h)
// }

// func (h MinHeap[T]) Less(i, j int) bool {
// 	return cmp.Less(h[i], h[j])
// }

// func (h MinHeap[T]) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func (h *MinHeap[T]) Pop() any {
// 	old := *h
// 	n := len(*h)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }

// func (h *MinHeap[T]) Push(x any) {
// 	*h = append(*h, x.(T))
// }

// func minMeetingRooms(intervals []Interval) int {
// 	n := len(intervals)
// 	if n <= 1 {
// 		return n
// 	}
// 	// for i := range n {
// 	// 	for j := i + 1; j < n; j++ {
// 	// 		if intervals[j].start < intervals[i].start {
// 	// 			intervals[i], intervals[j] = intervals[j], intervals[i]
// 	// 		}
// 	// 	}
// 	// }
// 	slices.SortFunc(intervals, func(a, b Interval) int {
// 		return cmp.Compare(a.start, b.start)
// 	})
// 	rooms := &MinHeap[int]{}
// 	heap.Init(rooms)
// 	for _, x := range intervals {
// 		if rooms.Len() > 0 && (*rooms)[0] <= x.start {
// 			heap.Pop(rooms)
// 		}
// 		heap.Push(rooms, x.end)
// 	}
// 	return rooms.Len()
// }

func minMeetingRooms(intervals []Interval) int {
	n := len(intervals)
	if n <= 1 {
		return n
	}
	starts, ends := make([]int, n), make([]int, n)
	for i, x := range intervals {
		starts[i] = x.start
		ends[i] = x.end
	}
	slices.Sort(starts)
	slices.Sort(ends)
	result, rooms := 0, 0
	i, j := 0, 0
	for i < len(intervals) {
		if starts[i] >= ends[j] {
			rooms--
			j++
		} else {
			rooms++
			i++
			if rooms > result {
				result = rooms
			}
		}
	}
	return result
}
