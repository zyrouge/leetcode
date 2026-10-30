package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", hammingWeight(23), 4)
	expect("#2", hammingWeight(2147483645), 30)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func hammingWeight(n int) int {
	result := 0
	c := n
	for c > 0 {
		if c&1 == 1 {
			result++
		}
		c >>= 1
	}
	return result
}
