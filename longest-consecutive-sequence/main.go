package main

import "fmt"

func main() {
	// nums := []int{2, 20, 4, 10, 3, 4, 5}
	nums := []int{0, 3, 2, 5, 4, 6, 1, 1}
	fmt.Printf("%d\n", longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, x := range nums {
		m[x] = true
	}
	result := 0
	for _, x := range nums {
		if _, ok := m[x-1]; !ok {
			c := 1
			for {
				if _, ok := m[x+c]; ok {
					c++
				} else {
					break
				}
			}
			if c > result {
				result = c
			}
		}
	}
	return result
}
