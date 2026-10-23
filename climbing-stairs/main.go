package main

import (
	"fmt"
)

func main() {
	expect("#1", climbStairs(0), 0)
	expect("#2", climbStairs(1), 1)
	expect("#3", climbStairs(2), 2)
	expect("#4", climbStairs(3), 3)
	expect("#5", climbStairs(45), 1836311903)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		panic(fmt.Sprintf("> fail: %s (expected: %v, got: %v)\n", message, expected, got))
	}
	fmt.Printf("> pass: %s (got: %v)\n", message, got)
}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	previous := 2
	current := 3
	for i := 3; i < n; i++ {
		temp := current
		current += previous
		previous = temp
	}
	return current
}
