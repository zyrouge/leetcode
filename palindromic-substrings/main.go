package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", countSubstrings("abc"), 3)
	expect("#2", countSubstrings("aaa"), 6)
	expect("#3", countSubstrings("xkjkqlajprjwefilxgpdpebieswu"), 30)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func countSubstrings(s string) int {
	// n := len(s)
	// t := make([][]bool, n)
	// for i := range n {
	// 	t[i] = make([]bool, n)
	// }
	result := 0
	// for i := n - 1; i >= 0; i-- {
	// 	for j := i; j < n; j++ {
	// 		if s[i] == s[j] && (j-i <= 2 || t[i+1][j-1]) {
	// 			t[i][j] = true
	// 			result++
	// 		}
	// 	}
	// }
	// for i := range n {
	// 	for j := 0; j <= i; j++ {
	// 		if s[i] == s[j] && (i-j <= 2 || t[i-1][j+1]) {
	// 			t[i][j] = true
	// 			result++
	// 		}
	// 	}
	// }
	for i := range s {
		result += countPalindromes(s, i, i)
		result += countPalindromes(s, i, i+1)
	}
	return result
}

func countPalindromes(s string, i, j int) int {
	result := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		result++
		i--
		j++
	}
	return result
}
