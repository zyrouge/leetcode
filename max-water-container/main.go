package main

import "fmt"

func main() {
	height := []int{1, 7, 2, 5, 4, 7, 3, 6}
	fmt.Printf("%d\n", maxArea(height))
}

func maxArea(heights []int) int {
	c := len(heights)
	result := 0
	i, j := 0, c-1
	for i < j {
		x, y := heights[i], heights[j]
		if area := (j - i) * min(x, y); area > result {
			result = area
		}
		if y < x {
			j--
			continue
		}
		i++
	}
	return result
}
