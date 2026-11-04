package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", longestCommonSubsequence("cat", "crabt"), 3)
	expect("#2", longestCommonSubsequence("abcd", "abcd"), 4)
	expect("#3", longestCommonSubsequence("abcd", "efgh"), 0)
	expect("#4", longestCommonSubsequence("bsbininm", "jmjkbkjkv"), 1)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// func longestCommonSubsequence(text1 string, text2 string) int {
// 	n1, n2 := len(text1), len(text2)
// 	m := make([][]int, n1)
// 	for i := range n1 {
// 		m[i] = make([]int, n2)
// 	}
// 	var dfs func(s1, s2 int) int
// 	dfs = func(s1, s2 int) int {
// 		if s1 >= n1 || s2 >= n2 {
// 			return 0
// 		}
// 		if m[s1][s2] != 0 {
// 			return m[s1][s2]
// 		}
// 		result := 0
// 		if text1[s1] == text2[s2] {
// 			result = 1 + dfs(s1+1, s2+1)
// 		} else {
// 			result = max(dfs(s1, s2+1), dfs(s1+1, s2))
// 		}
// 		m[s1][s2] = result
// 		return result
// 	}
// 	return dfs(0, 0)
// }

func longestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	m := make([][]int, n1+1)
	for i := range n1 + 1 {
		m[i] = make([]int, n2+1)
	}
	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				m[i][j] = 1 + m[i+1][j+1]
			} else {
				m[i][j] = max(m[i][j+1], m[i+1][j])
			}
		}
	}
	return m[0][0]
}
