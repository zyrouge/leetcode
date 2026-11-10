package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc(
		"#1",
		findWords([][]byte{
			{'a', 'b', 'c', 'd'},
			{'s', 'a', 'a', 't'},
			{'a', 'c', 'k', 'e'},
			{'a', 'c', 'd', 'n'}},
			[]string{"bat", "cat", "back", "backend", "stack"}),
		[]string{"cat", "back", "backend"},
		sortedSliceEquals,
	)
	expectFunc(
		"#2",
		findWords([][]byte{
			{'x', 'o'},
			{'x', 'o'}},
			[]string{"xoxo"}),
		[]string{},
		sortedSliceEquals,
	)
	expectFunc(
		"#3",
		findWords([][]byte{
			{'a', 'a'}},
			[]string{"a"}),
		[]string{"a"},
		sortedSliceEquals,
	)
	expectFunc(
		"#4",
		findWords([][]byte{
			{'a'}},
			[]string{"a"}),
		[]string{"a"},
		sortedSliceEquals,
	)
}

func sortedSliceEquals(a, b []string) bool {
	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

/* Attempt 1 */
// type BoardNode struct {
// 	Val    byte
// 	Left   *BoardNode
// 	Right  *BoardNode
// 	Top    *BoardNode
// 	Bottom *BoardNode
// }

// func findWords(board [][]byte, words []string) []string {
// 	m := NewBoardNodeMap(board)
// 	result := make([]string, 0)
// 	for _, word := range words {
// 		nodes := m[word[0]]
// 		for _, x := range nodes {
// 			// fmt.Println(string(word[0]), x)
// 			if x.Matches(word, 1, map[*BoardNode]bool{x: true}) {
// 				result = append(result, word)
// 				break
// 			}
// 		}
// 	}
// 	return result
// }

// func NewBoardNodeMap(board [][]byte) map[byte][]*BoardNode {
// 	xc, yc := len(board), len(board[0])
// 	m := make(map[byte][]*BoardNode)
// 	nodes := [][]*BoardNode{}
// 	for i := range xc {
// 		nodes = append(nodes, make([]*BoardNode, 0))
// 		for j := range yc {
// 			node := &BoardNode{board[i][j], nil, nil, nil, nil}
// 			if j > 0 {
// 				nodes[i][j-1].Right = node
// 				node.Left = nodes[i][j-1]
// 			}
// 			if i > 0 {
// 				nodes[i-1][j].Bottom = node
// 				node.Top = nodes[i-1][j]
// 			}
// 			mnodes := m[board[i][j]]
// 			if mnodes == nil {
// 				mnodes = make([]*BoardNode, 0)
// 			}
// 			m[board[i][j]] = append(mnodes, node)
// 			nodes[i] = append(nodes[i], node)
// 		}
// 	}
// 	// v := func(node *BoardNode) byte {
// 	// 	if node == nil {
// 	// 		return ' '
// 	// 	}
// 	// 	return node.Val
// 	// }
// 	// for i := range xc {
// 	// 	for j := range yc {
// 	// 		fmt.Printf("%c [%c|%c|%c|%c] ", v(nodes[i][j]), v(nodes[i][j].Top), v(nodes[i][j].Right), v(nodes[i][j].Bottom), v(nodes[i][j].Left))
// 	// 	}
// 	// 	fmt.Println()
// 	// }
// 	return m
// }

// // func (node *BoardNode) ValOrEmpty() byte {
// // 	if node == nil {
// // 		return ' '
// // 	}
// // 	return node.Val
// // }

// // func (node *BoardNode) Matches(word string) bool {
// // 	fmt.Println(word, string(node.Val))
// // 	var prev *BoardNode = nil
// // 	current := node
// // 	for i := 1; i < len(word); i++ {
// // 		x := word[i]
// // 		fmt.Printf("%c [%c|%c|%c|%c]\n", x, current.Top.ValOrEmpty(), current.Right.ValOrEmpty(), current.Bottom.ValOrEmpty(), current.Left.ValOrEmpty())
// // 		if current.Left != nil && current.Left != prev && current.Left.Val == x {
// // 			current = current.Left
// // 			continue
// // 		}
// // 		if current.Right != nil && current.Right != prev && current.Right.Val == x {
// // 			current = current.Right
// // 			continue
// // 		}
// // 		if current.Top != nil && current.Top != prev && current.Top.Val == x {
// // 			current = current.Top
// // 			continue
// // 		}
// // 		if current.Bottom != nil && current.Bottom != prev && current.Bottom.Val == x {
// // 			current = current.Bottom
// // 			continue
// // 		}
// // 		fmt.Println("no")
// // 		return false
// // 	}
// // 	return true
// // }

// func (node *BoardNode) Matches(word string, start int, visited map[*BoardNode]bool) bool {
// 	if start >= len(word) {
// 		return true
// 	}
// 	next := word[start]
// 	// fmt.Printf("%s %c [%c|%c|%c|%c]\n", word, next, node.Top.ValOrEmpty(), node.Right.ValOrEmpty(), node.Bottom.ValOrEmpty(), node.Left.ValOrEmpty())
// 	if node.Left != nil && !visited[node.Left] && node.Left.Val == next {
// 		visited[node.Left] = true
// 		if node.Left.Matches(word, start+1, visited) {
// 			return true
// 		}
// 		delete(visited, node.Left)
// 	}
// 	if node.Right != nil && !visited[node.Right] && node.Right.Val == next {
// 		visited[node.Right] = true
// 		if node.Right.Matches(word, start+1, visited) {
// 			return true
// 		}
// 		delete(visited, node.Right)
// 	}
// 	if node.Top != nil && !visited[node.Top] && node.Top.Val == next {
// 		visited[node.Top] = true
// 		if node.Top.Matches(word, start+1, visited) {
// 			return true
// 		}
// 		delete(visited, node.Top)
// 	}
// 	if node.Bottom != nil && !visited[node.Bottom] && node.Bottom.Val == next {
// 		visited[node.Bottom] = true
// 		if node.Bottom.Matches(word, start+1, visited) {
// 			return true
// 		}
// 		delete(visited, node.Bottom)
// 	}
// 	return false
// }

/* Attempt 2 */
type Trie struct {
	Val      byte
	Children [26]*Trie
	End      bool
	Count    int
}

func findWords(board [][]byte, words []string) []string {
	xn, yn := len(board), len(board[0])
	root := &Trie{}
	for _, x := range words {
		parent := root
		for _, c := range x {
			child := parent.Children[c-'a']
			if child == nil {
				child = &Trie{}
				parent.Children[c-'a'] = child
			}
			child.Count++
			parent = child
		}
		parent.End = true
	}
	result := make([]string, 0)
	var dfs func(i, j int, previousWord string, parent *Trie) int
	dfs = func(i, j int, previousWord string, parent *Trie) int {
		if i < 0 || i >= xn || j < 0 || j >= yn || board[i][j] == ' ' {
			return 0
		}
		letter := board[i][j]
		child := parent.Children[letter-'a']
		if child == nil || child.Count == 0 {
			return 0
		}
		found := 0
		word := previousWord + string(letter)
		if child.End {
			result = append(result, word)
			child.End = false
			found++
		}
		board[i][j] = ' '
		found += dfs(i-1, j, word, child)
		found += dfs(i+1, j, word, child)
		found += dfs(i, j-1, word, child)
		found += dfs(i, j+1, word, child)
		board[i][j] = letter
		child.Count -= found
		if child.Count == 0 {
			parent.Children[letter-'a'] = nil
		}
		return found
	}
	for i := range xn {
		for j := range yn {
			dfs(i, j, "", root)
		}
	}
	return result
}
