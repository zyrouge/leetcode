package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", numIslands([][]byte{
		{'0', '1', '1', '1', '0'},
		{'0', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}), 1)
	expect("#1", numIslands([][]byte{
		{'1', '1', '0', '0', '1'},
		{'1', '1', '0', '0', '1'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}), 4)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func numIslands(grid [][]byte) int {
	in, jn := len(grid), len(grid[0])
	result := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= in || j < 0 || j >= jn || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '0'
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := range in {
		for j := range jn {
			if grid[i][j] == '1' {
				result++
				dfs(i, j)
			}
		}
	}
	return result
}
