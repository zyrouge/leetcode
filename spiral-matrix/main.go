package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	// expectFunc("#1", spiralOrder([][]int{{1, 2}, {3, 4}}), []int{1, 2, 4, 3}, slices.Equal)
	expectFunc("#2", spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, slices.Equal)
	expectFunc("#3", spiralOrder([][]int{{7}, {9}, {6}}), []int{7, 9, 6}, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0)
	top, bottom, left, right := 0, m, 0, n
	for top < bottom && left < right {
		for i := left; i < right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		for i := top; i < bottom; i++ {
			result = append(result, matrix[i][right-1])
		}
		right--
		if !(left < right && top < bottom) {
			break
		}
		for i := right - 1; i >= left; i-- {
			result = append(result, matrix[bottom-1][i])
		}
		bottom--
		for i := bottom - 1; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
	}
	return result
}
