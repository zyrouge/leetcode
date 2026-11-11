package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 10), true)
	expect("#2", searchMatrix([][]int{{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40}}, 15), false)
	expect("#3", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3), true)
	expect("#4", searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 11), true)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// func searchMatrix(matrix [][]int, target int) bool {
// 	xn, yn := len(matrix), len(matrix[0])
// 	row := -1
// 	start, end := 0, xn-1
// 	for start <= end {
// 		row = start + ((end - start) / 2)
// 		if target > matrix[row][yn-1] {
// 			start = row + 1
// 		} else if target < matrix[row][0] {
// 			end = row - 1
// 		} else {
// 			break
// 		}
// 	}
// 	fmt.Println(row)
// 	if row == -1 {
// 		return false
// 	}
// 	start, end = 0, yn-1
// 	for start <= end {
// 		mid := start + ((end - start) / 2)
// 		if target > matrix[row][mid] {
// 			start = mid + 1
// 		} else if target < matrix[row][mid] {
// 			end = mid - 1
// 		} else {
// 			return true
// 		}
// 	}
// 	return false
// }

func searchMatrix(matrix [][]int, target int) bool {
	xn, yn := len(matrix), len(matrix[0])
	n := xn * yn
	start, end := 0, n-1
	for start <= end {
		mid := start + ((end - start) / 2)
		row, column := mid/yn, mid%yn
		if target > matrix[row][column] {
			start = mid + 1
		} else if target < matrix[row][column] {
			end = mid - 1
		} else {
			return true
		}
	}
	return false
}
