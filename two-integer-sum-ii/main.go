package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc("#1", twoSum([]int{1, 2, 3, 4}, 3), []int{1, 2}, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	i, j := 0, n-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i + 1, j + 1}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return nil
}
