package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc("#1", countBits(4), []int{0, 1, 1, 2, 1}, slices.Equal)
	expectFunc("#2", countBits(8), []int{0, 1, 1, 2, 1, 2, 2, 3, 1}, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func countBits(n int) []int {
	result := make([]int, n+1)
	current, next := 0, 2
	for i := 1; i <= n; i++ {
		if i == next {
			current = next
			next *= 2
		}
		result[i] = 1 + result[i-current]
	}
	return result
}
