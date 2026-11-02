package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", maxProduct([]int{2, 4, -3, 5}), 8)
	expect("#2", maxProduct([]int{-3, 0, -2}), 0)
	expect("#3", maxProduct([]int{-2}), -2)
	expect("#4", maxProduct([]int{0, 2}), 2)
	expect("#5", maxProduct([]int{3, -1, 4}), 4)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// func maxProduct(nums []int) int {
// 	n := len(nums)
// 	result := nums[0]
// 	for i := range n {
// 		current := nums[i]
// 		if current > result {
// 			result = current
// 		}
// 		for j := i + 1; j < n; j++ {
// 			current *= nums[j]
// 			if current > result {
// 				result = current
// 			}
// 		}
// 	}
// 	return result
// }

func maxProduct(nums []int) int {
	result := nums[0]
	cmin, cmax := 1, 1
	for _, x := range nums {
		temp := cmax * x
		cmax = max(x, temp, cmin*x)
		cmin = min(x, temp, cmin*x)
		result = max(result, cmax)
	}
	return result
}
