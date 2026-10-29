package main

import (
	"fmt"
	"os"
)

func main() {
	expect(
		"#1",
		exist([][]byte{
			{'A', 'B', 'C', 'D'},
			{'S', 'A', 'A', 'T'},
			{'A', 'C', 'A', 'E'}},
			"CAT"),
		true,
	)
	expect(
		"#2",
		exist([][]byte{
			{'A', 'B', 'C', 'D'},
			{'S', 'A', 'A', 'T'},
			{'A', 'C', 'A', 'E'}},
			"BAT"),
		false,
	)
	expect(
		"#3",
		exist([][]byte{
			{'a'}},
			"a"),
		true,
	)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func exist(board [][]byte, word string) bool {
	c := len(word)
	bxc, byc := len(board), len(board[0])
	visited := make([][]bool, bxc)
	for i := range bxc {
		visited[i] = make([]bool, byc)
	}
	var check func(start int, i int, j int) bool
	check = func(start int, i int, j int) bool {
		// fmt.Printf("%d %d:%d\n", start, i, j)
		if start == c {
			return true
		}
		if i < 0 || i >= bxc || j < 0 || j >= byc || board[i][j] != word[start] || visited[i][j] {
			return false
		}
		// fmt.Printf("%c\n", word[start])
		visited[i][j] = true
		found := check(start+1, i-1, j) || check(start+1, i+1, j) || check(start+1, i, j-1) || check(start+1, i, j+1)
		visited[i][j] = false
		return found
	}
	for i, x := range board {
		for j := range x {
			if check(0, i, j) {
				return true
			}
		}
	}
	return false
}
