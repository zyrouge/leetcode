package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", rob([]int{3, 4, 3}), 4)
	expect("#2", rob([]int{2, 9, 8, 3, 6}), 15)
	expect("#3", rob([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 21, 23, 25, 27, 29, 31, 33, 35, 37, 39, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 99, 97, 95, 93, 91, 89, 87, 85, 83, 81, 79, 77, 75, 73, 71, 69, 67, 65, 63, 61, 59, 57, 55, 53, 51, 49, 47, 45, 43, 41, 39, 37, 35, 33, 31, 29, 27, 25, 23}), 2540)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(nums[0], rob2(nums[1:]), rob2(nums[:n-1]))
}

func rob2(nums []int) int {
	result1, result2 := 0, 0
	for _, x := range nums {
		current := max(result1, result2+x)
		result2 = result1
		result1 = current
	}
	return result1
}
