package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", uniquePaths(3, 6), 21)
	expect("#2", uniquePaths(3, 3), 6)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func uniquePaths(m int, n int) int {
	cache := make([][]int, m)
	for i := range m {
		cache[i] = make([]int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if cache[i][j] != 0 {
			return cache[i][j]
		}
		if i == m-1 || j == n-1 {
			cache[i][j]++
			return cache[i][j]
		}
		cache[i][j] += dfs(i+1, j) + dfs(i, j+1)
		return cache[i][j]
	}
	dfs(0, 0)
	return cache[0][0]
}
