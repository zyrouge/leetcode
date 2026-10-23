package main

import "fmt"

type set[T comparable] map[T]bool

func main() {
	// board := [][]byte{
	// 	{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
	// 	{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '.', '3'},
	// 	{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
	// 	{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	board := [][]byte{{'1', '2', '.', '.', '3', '.', '.', '.', '.'},
		{'4', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '9', '1', '.', '.', '.', '.', '.', '3'},
		{'5', '.', '.', '.', '6', '.', '.', '.', '4'},
		{'.', '.', '.', '8', '.', '3', '.', '.', '5'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '.', '.', '.', '.', '.', '2', '.', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '8'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]set[byte], 9)
	cols := make(map[int]set[byte], 9)
	subs := make(map[int]set[byte], 9)
	for i := range 9 {
		rows[i] = make(set[byte], 9)
		cols[i] = make(set[byte], 9)
		subs[i] = make(set[byte], 9)
	}
	for i := range 9 {
		row := rows[i]
		for j := range 9 {
			target := board[j][i]
			fmt.Printf("%c (i: %d, j: %d) ", target, i, j)
			if target == byte('.') {
				fmt.Printf("\n")
				continue
			}
			if _, ok := row[target]; ok {
				return false
			}
			row[target] = true
			col := cols[j]
			if _, ok := col[target]; ok {
				return false
			}
			col[target] = true
			// [0, 0] = 0
			// [0, 2] = 0
			// [0, 3] = 1
			// [1, 3] = 1
			// [3, 0] = 3
			// [3, 3] = 4
			// [3, 7] = 5
			// [6, 0] = 6
			// [7, 0] = 6
			sub := subs[((i/3)*3)+(j/3)]
			fmt.Printf("(sub: %d (%d))\n", ((i/3)*3)+(j%3), (i/3)*3)
			if _, ok := sub[target]; ok {
				return false
			}
			sub[target] = true
		}
	}
	return true
}
