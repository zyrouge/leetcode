package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	expect("#1", foreignDictionary([]string{"z", "o"}), "zo")
	expect("#2", foreignDictionary([]string{"hrn", "hrf", "er", "enn", "rfnn"}), "hernf")
	expect("#3", foreignDictionary([]string{"abc", "ab"}), "")
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func foreignDictionary(words []string) string {
	n := len(words)
	adj := make(map[byte]map[byte]bool)
	for _, x := range words {
		for _, c := range x {
			if _, ok := adj[byte(c)]; !ok {
				adj[byte(c)] = make(map[byte]bool)
			}
		}
	}
	for i := 0; i < n-1; i++ {
		current, next := words[i], words[i+1]
		minLength := min(len(current), len(next))
		if len(current) > len(next) && next[:minLength] == current[:minLength] {
			return ""
		}
		for j := range minLength {
			if current[j] != next[j] {
				adj[current[j]][next[j]] = true
				break
			}
		}
	}
	ordered := make([]byte, 0)
	visited := make(map[byte]bool)
	var dfs func(x byte) bool
	dfs = func(x byte) bool {
		if current, ok := visited[x]; ok {
			return !current
		}
		visited[x] = true
		for y := range adj[x] {
			if !dfs(y) {
				return false
			}
		}
		visited[x] = false
		ordered = append(ordered, x)
		return true
	}
	for x := range adj {
		if !dfs(x) {
			return ""
		}
	}
	var result strings.Builder
	for _, x := range slices.Backward(ordered) {
		result.WriteString(string(x))
	}
	return result.String()
}
