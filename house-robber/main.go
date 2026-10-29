package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", rob([]int{1, 1, 3, 3}), 4)
	expect("#2", rob([]int{2, 9, 8, 3, 6}), 16)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func rob(nums []int) int {
	c := len(nums)
	m := make(map[int]int)
	var dfs func(i int) int
	dfs = func(i int) int {
		if i >= c {
			return 0
		}
		if total, ok := m[i]; ok {
			return total
		}
		m[i] = max(dfs(i+1), nums[i]+dfs(i+2))
		return m[i]
	}
	return dfs(0)
}
