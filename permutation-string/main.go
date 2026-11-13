package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", checkInclusion("abc", "lecabee"), true)
	expect("#2", checkInclusion("abc", "lecaabee"), false)
	expect("#3", checkInclusion("abc", "ccccbbbbaaaa"), false)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func checkInclusion(s1 string, s2 string) bool {
	s1n, s2n := len(s1), len(s2)
	if s1n > s2n {
		return false
	}
	s1c := [26]int{}
	for _, x := range s1 {
		s1c[x-'a']++
	}
	s2c := [26]int{}
	for i := range s1n {
		s2c[s2[i]-'a']++
	}
	if s1c == s2c {
		return true
	}
	for i := s1n; i < s2n; i++ {
		s2c[s2[i]-'a']++
		s2c[s2[i-s1n]-'a']--
		if s1c == s2c {
			return true
		}
	}
	return false
}
