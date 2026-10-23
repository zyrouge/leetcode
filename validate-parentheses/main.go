package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", isValid("[]"), true)
	expect("#2", isValid("([{}])"), true)
	expect("#3", isValid("[(])"), false)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func isValid(s string) bool {
	pairs := map[rune]rune{']': '[', ')': '(', '}': '{'}
	stack := make([]rune, len(s))
	tail := -1
	for _, x := range s {
		if m, ok := pairs[x]; ok {
			if tail == -1 || stack[tail] != m {
				return false
			}
			tail--
			continue
		}
		tail++
		stack[tail] = x
	}
	return tail == -1
}
