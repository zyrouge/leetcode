package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc("#1", rotated([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}), [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}, nestedSliceEqualFunc)
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

func rotated(matrix [][]int) [][]int {
	n := len(matrix)
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
		copy(result[i], matrix[i])
	}
	rotate(result)
	return result
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := range n / 2 {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
