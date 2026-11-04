package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}}), 2)
	expect("#2", countComponents(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}), 1)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// func countComponents(n int, edges [][]int) int {
// 	adjacents := make([][]int, n)
// 	for i := range n {
// 		adjacents[i] = make([]int, 0)
// 	}
// 	for _, x := range edges {
// 		adjacents[x[0]] = append(adjacents[x[0]], x[1])
// 		adjacents[x[1]] = append(adjacents[x[1]], x[0])
// 	}
// 	visited := make([]bool, n)
// 	var dfs func(x int)
// 	dfs = func(x int) {
// 		visited[x] = true
// 		for _, y := range adjacents[x] {
// 			if !visited[y] {
// 				dfs(y)
// 			}
// 		}
// 	}
// 	result := 0
// 	for x := range n {
// 		if !visited[x] {
// 			dfs(x)
// 			result++
// 		}
// 	}
// 	return result
// }

func countComponents(n int, edges [][]int) int {
	adjacents := make([][]int, n)
	for i := range n {
		adjacents[i] = make([]int, 0)
	}
	for _, x := range edges {
		adjacents[x[0]] = append(adjacents[x[0]], x[1])
		adjacents[x[1]] = append(adjacents[x[1]], x[0])
	}
	visited := make([]bool, n)
	var bfs func(x int)
	bfs = func(x int) {
		queue := []int{x}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if visited[node] {
				continue
			}
			visited[node] = true
			queue = append(queue, adjacents[node]...)
		}
	}
	result := 0
	for x := range n {
		if !visited[x] {
			bfs(x)
			result++
		}
	}
	return result
}
