package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc("#1", combinationSum([]int{2, 5, 6, 9}, 9), [][]int{{2, 2, 5}, {9}}, nestedSlicesSortedEqual)
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

func combinationSum(nums []int, target int) [][]int {
	c := len(nums)
	for i := range c {
		for j := i + 1; j < c; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	// slices.Sort(nums)
	result := make([][]int, 0)
	var backtrack func(start int, current []int, currentSum int)
	backtrack = func(start int, current []int, total int) {
		// fmt.Println(start, current, total)
		if total == target {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}
		if start >= c || total > target {
			return
		}
		current = append(current, nums[start])
		backtrack(start, current, total+nums[start])
		current = current[:len(current)-1]
		backtrack(start+1, current, total)
	}
	backtrack(0, make([]int, 0), 0)
	return result
}
