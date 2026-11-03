package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc("#1", pacificAtlantic([][]int{
		{4, 2, 7, 3, 4},
		{7, 4, 6, 4, 7},
		{6, 3, 5, 3, 6},
	}), [][]int{{0, 2}, {0, 4}, {1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {2, 0}}, nestedSlicesSortedEqual)
	expectFunc("#2", pacificAtlantic([][]int{{1}, {1}}), [][]int{{0, 0}, {1, 0}}, nestedSlicesSortedEqual)
	expectFunc("#3", pacificAtlantic([][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}),
		[][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 9}, {0, 10}, {0, 11},
			{0, 12}, {0, 13}, {0, 14}, {0, 15}, {0, 16}, {0, 17}, {0, 18}, {0, 19}, {1, 0}, {1, 19}, {2, 0}, {2, 19},
			{3, 0}, {3, 19}, {4, 0}, {4, 19}, {5, 0}, {5, 19}, {6, 0}, {6, 19}, {7, 0}, {7, 19}, {8, 0}, {8, 19},
			{9, 0}, {9, 19}, {10, 0}, {10, 19}, {11, 0}, {11, 19}, {12, 0}, {12, 19}, {13, 0}, {13, 19}, {14, 0}, {14, 19},
			{15, 0}, {15, 19}, {16, 0}, {16, 19}, {17, 0}, {17, 19}, {18, 0}, {18, 19}, {19, 0}, {19, 1}, {19, 2},
			{19, 3}, {19, 4}, {19, 5}, {19, 6}, {19, 7}, {19, 8}, {19, 9}, {19, 10}, {19, 11}, {19, 12}, {19, 13}, {19, 14},
			{19, 15}, {19, 16}, {19, 17}, {19, 18}, {19, 19}}, nestedSlicesSortedEqual)
	expectFunc("#4", pacificAtlantic([][]int{
		{1, 1, 1},
		{1, 1, 1},
		{5, 1, 5},
	}), [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}, nestedSlicesSortedEqual)
}

func nestedSlicesSortedEqual[T cmp.Ordered](a, b [][]T) bool {
	for _, x := range a {
		slices.Sort(x)
	}
	for _, x := range b {
		slices.Sort(x)
	}
	slices.SortFunc(a, slices.Compare)
	slices.SortFunc(b, slices.Compare)
	return slices.EqualFunc(a, b, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func pacificAtlantic(heights [][]int) [][]int {
	in, jn := len(heights), len(heights[0])
	m := make([][][4]*[2]bool, in)
	used := make([][]bool, in)
	for i := range in {
		m[i] = make([][4]*[2]bool, jn)
		used[i] = make([]bool, jn)
	}
	var dfs func(pi, pj, pd int, i, j int) [2]bool
	dfs = func(pi, pj, pd int, i, j int) [2]bool {
		if i < 0 || j < 0 {
			return [2]bool{true, false}
		}
		if i >= in || j >= jn {
			return [2]bool{false, true}
		}
		if pd != -1 {
			if found := m[i][j][pd]; found != nil {
				return *found
			}
		}
		if heights[i][j] > heights[pi][pj] || used[i][j] {
			return [2]bool{false, false}
		}
		// fmt.Printf("   %d i=%d j=%d\n", heights[i][j], i, j)
		used[i][j] = true
		result := dfs(i, j, 0, i-1, j) // top
		if !result[0] || !result[1] {
			temp := dfs(i, j, 1, i, j-1) // left
			result[0] = temp[0] || result[0]
			result[1] = temp[1] || result[1]
		}
		if !result[0] || !result[1] {
			temp := dfs(i, j, 2, i+1, j) // bottom
			result[0] = temp[0] || result[0]
			result[1] = temp[1] || result[1]
		}
		if !result[0] || !result[1] {
			temp := dfs(i, j, 3, i, j+1) // right
			result[0] = temp[0] || result[0]
			result[1] = temp[1] || result[1]
		}
		used[i][j] = false
		if pd != -1 {
			m[i][j][pd] = &result
		}
		return result
	}
	result := make([][]int, 0)
	for i, x := range heights {
		for j := range x {
			// fmt.Printf("->\n")
			reach := dfs(i, j, -1, i, j)
			if reach[0] && reach[1] {
				result = append(result, []int{i, j})
			}
			// fmt.Printf("<- pacific=%v atlantic=%v\n", reach[0], reach[1])
		}
	}
	return result
}
