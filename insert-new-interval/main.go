package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc("#1", insert([][]int{{1, 3}, {4, 6}}, []int{2, 5}), [][]int{{1, 6}}, nestedSliceEqualFunc)
	expectFunc("#2", insert([][]int{{1, 2}, {3, 5}, {9, 10}}, []int{6, 7}), [][]int{{1, 2}, {3, 5}, {6, 7}, {9, 10}}, nestedSliceEqualFunc)
}

func nestedSliceEqualFunc[T comparable](a, b [][]T) bool {
	return slices.EqualFunc(a, b, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0)
	for i, x := range intervals {
		if newInterval[1] < x[0] {
			result = append(result, newInterval)
			return append(result, intervals[i:]...)
		} else if newInterval[0] > x[1] {
			result = append(result, x)
		} else {
			newInterval[0] = min(x[0], newInterval[0])
			newInterval[1] = max(x[1], newInterval[1])
		}
	}
	result = append(result, newInterval)
	return result
}
