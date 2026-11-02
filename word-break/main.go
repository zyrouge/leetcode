package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", wordBreak("neetcode", []string{"neet", "code"}), true)
	expect("#2", wordBreak("applepenapple", []string{"apple", "pen", "ape"}), true)
	expect("#3", wordBreak("catsincars", []string{"cats", "cat", "sin", "in", "car"}), false)
	expect("#4", wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa", "ba"}), false)
	expect("#5", wordBreak("a", []string{"a"}), true)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	m := make([]int, n)
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i >= n {
			return true
		}
		if found := m[i]; found != 0 {
			return found == 1
		}
		for _, x := range wordDict {
			if i+len(x) <= n && s[i:i+len(x)] == x && dfs(i+len(x)) {
				m[i] = 1
				return true
			}
		}
		m[i] = -1
		return false
	}
	return dfs(0)
}

// type TrieNode struct {
// 	Children  [26]*TrieNode
// 	EndOfWord bool
// }

// func wordBreak(s string, wordDict []string) bool {
// 	children := NewTrieNode(wordDict)
// 	m := make([]int, len(s))
// 	return children.WordBreak(m, s, 0)
// }

// func NewTrieNode(wordDict []string) *TrieNode {
// 	root := &TrieNode{}
// 	for _, word := range wordDict {
// 		node := root
// 		for i := 0; i < len(word); i++ {
// 			x := word[i]
// 			child := node.Children[x-'a']
// 			if child == nil {
// 				child = &TrieNode{[26]*TrieNode{}, false}
// 				node.Children[x-'a'] = child
// 			}
// 			node = child
// 		}
// 		node.EndOfWord = true
// 	}
// 	return root
// }

// func (root *TrieNode) WordBreak(m []int, s string, start int) bool {
// 	n := len(s)
// 	if start == n {
// 		return true
// 	}
// 	found := m[start]
// 	if found == -1 {
// 		return false
// 	}
// 	if found != 0 {
// 		return root.WordBreak(m, s, found)
// 	}
// 	node, i := root, start
// 	for i < n {
// 		child := node.Children[s[i]-'a']
// 		if child == nil {
// 			break
// 		}
// 		i++
// 		if child.EndOfWord && root.WordBreak(m, s, i) {
// 			m[start] = n
// 			return true
// 		}
// 		node = child
// 	}
// 	m[start] = -1
// 	return false
// }
