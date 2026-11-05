package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", canJump([]int{1, 2, 0, 1, 0}), true)
	expect("#2", canJump([]int{1, 2, 1, 0, 1}), false)
	expect("#3", canJump([]int{2, 0}), true)
	expect("#4", canJump([]int{2, 5, 0, 0}), true)
	expect("#5", canJump([]int{3, 2, 1, 0, 4}), false)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// func canJump(nums []int) bool {
// 	n := len(nums)
// 	var dfs func(i, jumps int) bool
// 	dfs = func(i, jumps int) bool {
// 		if jumps < 0 {
// 			return false
// 		}
// 		if i >= n-1 {
// 			return true
// 		}
// 		return dfs(i+1, max(nums[i], jumps)-1)
// 	}
// 	return dfs(0, nums[0])
// }

func canJump(nums []int) bool {
	n := len(nums)
	i, jumps := 0, 0
	for i < n {
		jumps = max(nums[i], jumps) - 1
		i++
		if jumps < 0 {
			break
		}
	}
	return i == n
}
