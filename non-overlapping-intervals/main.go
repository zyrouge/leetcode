package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	expect("#1", eraseOverlapIntervals([][]int{{1, 2}, {2, 4}, {1, 4}}), 1)
	expect("#2", eraseOverlapIntervals([][]int{{1, 2}, {2, 4}}), 0)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	// for i := range intervals {
	// 	for j := i + 1; j < n; j++ {
	// 		if intervals[j][1] < intervals[i][1] {
	// 			intervals[i], intervals[j] = intervals[j], intervals[i]
	// 		}
	// 	}
	// }
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})
	result := 0
	lastEnd := intervals[0][1]
	for i := 1; i < n; i++ {
		if lastEnd > intervals[i][0] {
			result++
		} else {
			lastEnd = intervals[i][1]
		}
	}
	return result
}
