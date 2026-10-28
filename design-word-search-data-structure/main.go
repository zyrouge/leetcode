package main

import (
	"fmt"
	"os"
)

func main() {
	{
		wordDictionary := Constructor()
		wordDictionary.AddWord("day")
		wordDictionary.AddWord("bay")
		wordDictionary.AddWord("may")
		expect("#1 > #1", wordDictionary.Search("say"), false)
		expect("#1 > #2", wordDictionary.Search("day"), true)
		expect("#1 > #3", wordDictionary.Search(".ay"), true)
		expect("#1 > #4", wordDictionary.Search("b.."), true)
		expect("#1 > #5", wordDictionary.Search(".."), false)
	}
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

type WordDictionary struct {
	Root *WordDictionaryNode
}

type WordDictionaryNode struct {
	Nodes     [26]*WordDictionaryNode
	EndOfWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{&WordDictionaryNode{[26]*WordDictionaryNode{}, false}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.Root
	for _, x := range word {
		temp := node.Nodes[x-'a']
		if temp == nil {
			temp = &WordDictionaryNode{[26]*WordDictionaryNode{}, false}
			node.Nodes[x-'a'] = temp
		}
		node = temp
	}
	node.EndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.Root.Search(word, 0)
}

// func (this *WordDictionaryNode) Search(word string, start int) bool {
// 	if start >= len(word) {
// 		return this.EndOfWord
// 	}
// 	letter := word[start]
// 	start++
// 	if letter == '.' {
// 		for i := range 26 {
// 			if x := this.Nodes[i]; x != nil && x.Search(word, start) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// 	temp := this.Nodes[letter-'a']
// 	return temp != nil && temp.Search(word, start)
// }

func (this *WordDictionaryNode) Search(word string, start int) bool {
	current := this
	for i := start; i < len(word); i++ {
		x := word[i]
		if x == '.' {
			for _, y := range current.Nodes {
				if y != nil && y.Search(word, i+1) {
					return true
				}
			}
			return false
		}
		current = current.Nodes[x-'a']
		if current == nil {
			return false
		}
	}
	return current.EndOfWord
}
