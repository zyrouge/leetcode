package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}), true)
	expect("#2", validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}), false)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func validTree(n int, edges [][]int) bool {
	adjacents := make([][]int, n)
	for i := range n {
		adjacents[i] = make([]int, 0)
	}
	for _, x := range edges {
		adjacents[x[0]] = append(adjacents[x[0]], x[1])
		adjacents[x[1]] = append(adjacents[x[1]], x[0])
	}
	visited := make([]bool, n)
	c := 0
	var dfs func(x, p int) bool
	dfs = func(x, p int) bool {
		if visited[x] {
			return false
		}
		visited[x] = true
		c++
		for _, y := range adjacents[x] {
			if y == p {
				continue
			}
			if !dfs(y, x) {
				return false
			}
		}
		return true
	}
	return dfs(0, -1) && c == n
}
