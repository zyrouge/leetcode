package main

import (
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	expect("#1", characterReplacement("XYYX", 2), 4)
	expect("#2", characterReplacement("AAABABB", 1), 5)
	expect("#3", characterReplacement("ABBB", 2), 4)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func characterReplacement(s string, k int) int {
	c := len(s)
	if c == 0 {
		return 0
	}
	result := 0
	counts := make(map[byte]int)
	start := 0
	end := 0
	current := 0
	for end < c {
		x := s[end]
		counts[x]++
		if counts[x] > current {
			current = counts[x]
		}
		for end-start-current+1 > k {
			counts[s[start]]--
			start++
		}
		result = max(result, end-start+1)
		// fmt.Printf("start=[%d]%c end=[%d]%c counts=%v result=%d\n", start, s[start], end, s[end], counts, result)
		end++
	}
	return result
}
