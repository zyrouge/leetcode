package main

import (
	"fmt"
	"os"
)

func main() {
	{
		prefixTree := Constructor()
		prefixTree.Insert("dog")
		expect("#1 > #1", prefixTree.Search("dog"), true)
		expect("#1 > #2", prefixTree.Search("do"), false)
		expect("#1 > #3", prefixTree.StartsWith("do"), true)
		prefixTree.Insert("do")
		expect("#1 > #4", prefixTree.Search("do"), true)
	}
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

type PrefixTree struct {
	Root *PrefixTreeNode
}

type PrefixTreeNode struct {
	Nodes     [26]*PrefixTreeNode
	EndOfWord bool
}

func Constructor() PrefixTree {
	return PrefixTree{&PrefixTreeNode{[26]*PrefixTreeNode{}, false}}
}

func (this *PrefixTree) Insert(word string) {
	node := this.Root
	for _, x := range word {
		temp := node.Nodes[x-'a']
		if temp == nil {
			temp = &PrefixTreeNode{[26]*PrefixTreeNode{}, false}
			node.Nodes[x-'a'] = temp
		}
		node = temp
	}
	node.EndOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	node := this.Root
	for _, x := range word {
		node = node.Nodes[x-'a']
		if node == nil {
			return false
		}
	}
	return node.EndOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this.Root
	for _, x := range prefix {
		node = node.Nodes[x-'a']
		if node == nil {
			return false
		}
	}
	return true
}
