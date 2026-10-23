package main

import (
	"fmt"
)

func main() {
	expect("#1", lengthOfLongestSubstring("zxyzxyz"), 3)
	expect("#2", lengthOfLongestSubstring("xxxx"), 1)
	expect("#3", lengthOfLongestSubstring(" "), 1)
	expect("#4", lengthOfLongestSubstring(""), 0)
	expect("#5", lengthOfLongestSubstring("abcabcbb"), 3)
	expect("#6", lengthOfLongestSubstring("au"), 2)
	expect("#7", lengthOfLongestSubstring("abba"), 2)
	expect("#8", lengthOfLongestSubstring("thequickbrownfoxjumpsoverthelazydogthequickbrownfoxjumpsovert"), 17)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		panic(fmt.Sprintf("> fail: %s (expected: %v, got: %v)\n", message, expected, got))
	}
	fmt.Printf("> pass: %s (got: %v)\n", message, got)
}

func lengthOfLongestSubstring(s string) int {
	c := len(s)
	if c < 2 {
		return c
	}
	result := 0
	start := 0
	set := make(map[byte]int)
	for end := range c {
		x := s[end]
		at, ok := set[x]
		if ok {
			start = max(start, at+1)
		}
		set[x] = end
		result = max(result, end-start+1)
	}
	return result
}
