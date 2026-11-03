package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", canFinish(2, [][]int{{0, 1}}), true)
	expect("#2", canFinish(2, [][]int{{0, 1}, {1, 0}}), false)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	prerequisitesMap := make([][]int, numCourses)
	for i := range numCourses {
		prerequisitesMap[i] = make([]int, 0)
	}
	for _, x := range prerequisites {
		prerequisitesMap[x[0]] = append(prerequisitesMap[x[0]], x[1])
	}
	m := make([]bool, numCourses)
	visiting := make([]bool, numCourses)
	var dfs func(x int) bool
	dfs = func(x int) bool {
		if ok := m[x]; ok {
			return ok
		}
		if visiting[x] {
			return false
		}
		visiting[x] = true
		for _, y := range prerequisitesMap[x] {
			if !dfs(y) {
				return false
			}
		}
		visiting[x] = false
		m[x] = true
		return true
	}
	for x := range numCourses {
		if !dfs(x) {
			return false
		}
	}
	return true
}
