package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", findMin([]int{3, 4, 5, 6, 1, 2}), 1)
	expect("#2", findMin([]int{4, 5, 0, 1, 2, 3}), 0)
	expect("#3", findMin([]int{4, 5, 6, 7}), 4)
	expect("#4", findMin([]int{1, 2}), 1)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// brute-force
// func findMin(nums []int) int {
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] < nums[i-1] {
// 			return nums[i]
// 		}
// 	}
// 	return nums[0]
// }

func findMin(nums []int) int {
	c := len(nums)
	if c == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + ((right - left) / 2)
		// fmt.Printf("L[%d]%d M[%d]%d R[%d]%d\n", left, nums[left], mid, nums[mid], right, nums[right])
		if nums[mid] > nums[right] {
			left = mid + 1
			continue
		}
		right = mid
	}
	// fmt.Printf("L[%d]%d R[%d]%d\n", left, nums[left], right, nums[right])
	return nums[left]
}
