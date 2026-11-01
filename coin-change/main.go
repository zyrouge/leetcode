package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	expect("#1", coinChange([]int{1, 5, 10}, 12), 3)
	expect("#2", coinChange([]int{2}, 3), -1)
	expect("#3", coinChange([]int{1}, 0), 0)
	expect("#4", coinChange([]int{1, 2147483647}, 2), 2)
	expect("#5", coinChange([]int{1, 3, 4}, 6), 2)
	expect("#6", coinChange([]int{186, 419, 83, 408}, 6249), 20)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	m := make(map[int]int)
	var dfs func(remaining int) int
	dfs = func(remaining int) int {
		if found, ok := m[remaining]; ok {
			return found
		}
		if remaining == 0 {
			return 0
		}
		result := math.MaxInt32
		for _, x := range coins {
			if remaining-x >= 0 {
				result = min(result, 1+dfs(remaining-x))
			}
		}
		m[remaining] = result
		return result
	}
	result := dfs(amount)
	if result == math.MaxInt32 {
		return -1
	}
	return result
}
