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
