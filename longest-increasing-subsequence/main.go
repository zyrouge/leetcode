package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", lengthOfLIS([]int{9, 1, 4, 2, 3, 3, 7}), 4)
	expect("#2", lengthOfLIS([]int{0, 3, 1, 3, 2, 3}), 4)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	m := make([]int, n)
	var dfs func(start int) int
	dfs = func(start int) int {
		if found := m[start]; found != 0 {
			return found
		}
		result := 1
		for i := start + 1; i < n; i++ {
			if nums[i] > nums[start] {
				result = max(result, 1+dfs(i))
			}
		}
		m[start] = result
		return result
	}
	result := 0
	for i := range n {
		result = max(result, dfs(i))
	}
	return result
}
