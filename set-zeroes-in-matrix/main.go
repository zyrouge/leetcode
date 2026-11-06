package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc("#1", setZeroesCopied([][]int{{0, 1}, {1, 0}}), [][]int{{0, 0}, {0, 0}}, nestedSliceEqualFunc)
	expectFunc("#2", setZeroesCopied([][]int{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}}), [][]int{{1, 0, 3}, {0, 0, 0}, {6, 0, 8}}, nestedSliceEqualFunc)
	expectFunc("#3", setZeroesCopied([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}), [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}, nestedSliceEqualFunc)
}

func nestedSliceEqualFunc[T comparable](a, b [][]T) bool {
	return slices.EqualFunc(a, b, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func setZeroesCopied(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
		copy(result[i], matrix[i])
	}
	setZeroes(result)
	return result
}

// func setZeroes(matrix [][]int) {
// 	m, n := len(matrix), len(matrix[0])
// 	zeroColumns := make([]int, 0)
// 	for i := range m {
// 		zeroRow := false
// 		for j := range n {
// 			if matrix[i][j] != 0 {
// 				continue
// 			}
// 			zeroRow = true
// 			break
// 		}
// 		if zeroRow {
// 			for j := range n {
// 				if matrix[i][j] == 0 {
// 					zeroColumns = append(zeroColumns, j)
// 				}
// 				matrix[i][j] = 0
// 			}
// 		}
// 	}
// 	for _, j := range zeroColumns {
// 		for i := range m {
// 			matrix[i][j] = 0
// 		}
// 	}
// }

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	zeroRows, zeroColumns := make([]bool, m), make([]bool, n)
	for i, x := range matrix {
		for j, y := range x {
			if y == 0 {
				zeroRows[i] = true
				zeroColumns[j] = true
			}
		}
	}
	for i, x := range matrix {
		for j := range x {
			if zeroRows[i] || zeroColumns[j] {
				matrix[i][j] = 0
			}
		}
	}
}
