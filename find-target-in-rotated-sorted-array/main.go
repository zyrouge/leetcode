package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", search([]int{3, 4, 5, 6, 1, 2}, 1), 4)
	expect("#2", search([]int{3, 5, 6, 0, 1, 2}, 4), -1)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + ((right - left) / 2)
		fmt.Printf("L[%d]%d M[%d]%d R[%d]%d\n", left, nums[left], mid, nums[mid], right, nums[right])
		isLeft := false
		if nums[left] < nums[mid] {
			isLeft = target >= nums[left] && target <= nums[mid]
		} else {
			isLeft = target < nums[mid+1] || target > nums[right]
			// fmt.Println(target, nums[mid+1], nums[right], isLeft)
		}
		if isLeft {
			right = mid
			// fmt.Println("FIRST!")
			continue
		}
		left = mid + 1
	}
	// fmt.Printf("L[%d]%d R[%d]%d\n", left, nums[left], right, nums[right])
	if nums[left] == target {
		return left
	}
	return -1
}
