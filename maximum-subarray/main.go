package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", maxSubArray([]int{2, -3, 4, -2, 2, 1, -1, 4}), 8)
	expect("#2", maxSubArray([]int{-1}), -1)
	expect("#3", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// func maxSubArray(nums []int) int {
// 	n := len(nums)
// 	result := 0
// 	var dfs func(i int) int
// 	dfs = func(i int) int {
// 		if i == n-1 {
// 			result = nums[i]
// 			return nums[i]
// 		}
// 		currentResult := max(nums[i], nums[i]+dfs(i+1))
// 		result = max(result, currentResult)
// 		// fmt.Printf("i=%d num=%d currentResult=%d result=%d\n", i, nums[i], currentResult, result)
// 		return currentResult
// 	}
// 	dfs(0)
// 	return result
// }

func maxSubArray(nums []int) int {
	n := len(nums)
	result := nums[n-1]
	currentResult := result
	for i := n - 2; i >= 0; i-- {
		currentResult = max(nums[i], nums[i]+currentResult)
		result = max(result, currentResult)
	}
	return result
}
