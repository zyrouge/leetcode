package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", search([]int{-1, 0, 2, 4, 6, 8}, 4), 3)
	expect("#2", search([]int{-1, 0, 2, 4, 6, 8}, 3), -1)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// func search(nums []int, target int) int {
// 	n := len(nums)
// 	start, end := 0, n
// 	for end-start > 1 {
// 		mid := start + ((end - start) / 2)
// 		if target < nums[mid] {
// 			end = mid
// 		} else if target > nums[mid] {
// 			start = mid + 1
// 		} else {
// 			return mid
// 		}
// 	}
// 	if target == nums[start] {
// 		return start
// 	}
// 	return -1
// }

func search(nums []int, target int) int {
	n := len(nums)
	start, end := 0, n-1
	for start <= end {
		mid := start + ((end - start) / 2)
		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
