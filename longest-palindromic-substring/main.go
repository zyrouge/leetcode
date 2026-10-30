package main

import (
	"fmt"
	"os"
)

func main() {
	// expect("#1", longestPalindrome("ababd"), "aba")
	expect("#1", longestPalindrome("ababd"), "bab")
	expect("#2", longestPalindrome("abbc"), "bb")
	expect("#3", longestPalindrome("aacabdkacaa"), "aca")
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func longestPalindrome(s string) string {
	n := len(s)
	c := make([][]bool, n)
	for i := range n {
		c[i] = make([]bool, n)
	}
	start, startLen := 0, 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] && (j-i <= 2 || c[i+1][j-1]) {
				c[i][j] = true
				if j-i+1 > startLen {
					start = i
					startLen = j - i + 1
				}
			}
		}
	}
	return s[start : start+startLen]
}
