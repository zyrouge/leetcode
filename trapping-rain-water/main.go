package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", trap([]int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}), 9)
	expect("#2", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), 6)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func trap(height []int) int {
	n := len(height)
	result := 0
	left, right := 0, n-1
	leftMax, rightMax := height[left], height[right]
	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(leftMax, height[left])
			result += leftMax - height[left]
		} else {
			right--
			rightMax = max(rightMax, height[right])
			result += rightMax - height[right]
		}
	}
	return result
}
