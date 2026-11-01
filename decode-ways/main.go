package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", numDecodings("12"), 2)
	expect("#2", numDecodings("01"), 0)
	expect("#3", numDecodings("226"), 3)
	expect("#4", numDecodings("1212121212121212121212121212121212121212121277777777777777777777777777777777777777777777777777777777"), 1134903170)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func numDecodings(s string) int {
	n := len(s)
	m := make(map[int]int)
	var dfs func(i int) int
	dfs = func(i int) int {
		if found, ok := m[i]; ok {
			return found
		}
		if i == n {
			m[i] = 1
			return 1
		}
		if s[i] == '0' {
			m[i] = 0
			return 0
		}
		result := dfs(i + 1)
		if i+1 < n && (s[i] == '1' || (s[i] == '2' && (s[i+1] >= '0' && s[i+1] <= '6'))) {
			result += dfs(i + 2)
		}
		m[i] = result
		return result
	}
	return dfs(0)
}
