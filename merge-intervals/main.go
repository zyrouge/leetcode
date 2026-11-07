package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc("#1", merge([][]int{{1, 3}, {1, 5}, {6, 7}}), [][]int{{1, 5}, {6, 7}}, nestedSliceEqualFunc)
	expectFunc("#2", merge([][]int{{1, 2}, {2, 3}}), [][]int{{1, 3}}, nestedSliceEqualFunc)
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

// func merge(intervals [][]int) [][]int {
// 	n := len(intervals)
// 	for i := range n {
// 		for j := i + 1; j < n; j++ {
// 			if intervals[j][0] < intervals[i][0] {
// 				intervals[i], intervals[j] = intervals[j], intervals[i]
// 			}
// 		}
// 	}
// 	result := [][]int{intervals[0]}
// 	for i := 1; i < n; i++ {
// 		lastResult := result[len(result)-1]
// 		if intervals[i][0] <= lastResult[1] {
// 			if intervals[i][1] > lastResult[1] {
// 				lastResult[1] = intervals[i][1]
// 			}
// 			continue
// 		}
// 		result = append(result, intervals[i])
// 	}
// 	return result
// }

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	for i := range n {
		for j := i + 1; j < n; j++ {
			if intervals[j][0] < intervals[i][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	result := [][]int{}
	lastResult := intervals[0]
	for i := 1; i < n; i++ {
		if lastResult[1] < intervals[i][0] {
			result = append(result, lastResult)
			lastResult = intervals[i]
		} else {
			lastResult[1] = max(lastResult[1], intervals[i][1])
		}
	}
	return append(result, lastResult)
}
