package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", reverseBits(21), 2818572288)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func reverseBits(n int) int {
	result := 0
	for range 32 {
		result <<= 1
		result |= n & 1
		n >>= 1
	}
	return result
}
