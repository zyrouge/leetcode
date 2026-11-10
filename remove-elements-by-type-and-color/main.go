package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", canRemoveAll([]int{0, 1, 1, 0}, []int{5, 3, 3, 5}, []int{2, 2}), true)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func canRemoveAll(types []int, colors []int, groupSize []int) bool {
	groups := make([][2]int, 0)
	for i, t := range types {
		gn := len(groups)
		if gn > 0 {
			pt := groups[gn-1][0]
			if types[pt] == t && colors[pt] == colors[i] {
				groups[gn-1][1]++
				if groupSize[t] == (groups[gn-1][1] - groups[gn-1][0]) {
					groups = groups[:gn-1]
				}
				continue
			}
		}
		if groupSize[t] != 1 {
			groups = append(groups, [2]int{i, i + 1})
		}
	}
	return len(groups) == 0
}
