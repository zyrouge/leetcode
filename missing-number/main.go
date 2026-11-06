package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", missingNumber([]int{1, 2, 3}), 0)
	expect("#2", missingNumber([]int{0, 2}), 1)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// func missingNumber(nums []int) int {
// 	n := len(nums)
// 	m := make(map[int]bool)
// 	for x := range n {
// 		m[x] = true
// 	}
// 	for _, x := range nums {
// 		delete(m, x)
// 	}
// 	next, stop := iter.Pull2(maps.All(m))
// 	defer stop()
// 	if x, _, ok := next(); ok {
// 		return x
// 	}
// 	return n
// }

// func missingNumber(nums []int) int {
// 	n := len(nums)
// 	m := make([]bool, n+1)
// 	for _, x := range nums {
// 		m[x] = true
// 	}
// 	for i, x := range m {
// 		if !x {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func missingNumber(nums []int) int {
// 	n := len(nums)
// 	result := n * (n + 1) / 2
// 	for _, x := range nums {
// 		result -= x
// 	}
// 	return result
// }

func missingNumber(nums []int) int {
	n := len(nums)
	result := n
	for i, x := range nums {
		result += i - x
	}
	return result
}
