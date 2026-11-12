package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	expect("#1", minEatingSpeed([]int{1, 4, 3, 2}, 9), 2)
	expect("#2", minEatingSpeed([]int{25, 10, 23, 4}, 4), 25)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

// func minEatingSpeed(piles []int, h int) int {
// 	n := len(piles)
// 	for i := range piles {
// 		for j := i + 1; j < n; j++ {
// 			if piles[j] < piles[i] {
// 				piles[i], piles[j] = piles[j], piles[i]
// 			}
// 		}
// 	}
// 	result := piles[n-1]
// 	for {
// 		taken := 0
// 		for _, x := range piles {
// 			taken += int(math.Ceil(float64(x) / float64(result)))
// 		}
// 		if taken > h {
// 			result++
// 			break
// 		}
// 		result--
// 	}
// 	return result
// }

func minEatingSpeed(piles []int, h int) int {
	r := 0
	for _, x := range piles {
		if x > r {
			r = x
		}
	}
	l := 1
	result := -1
	for l <= r {
		current := l + ((r - l) / 2)
		taken := 0
		for _, x := range piles {
			taken += int(math.Ceil(float64(x) / float64(current)))
		}
		if taken <= h {
			result = current
			r = current - 1
		} else {
			l = current + 1
		}
	}
	return result
}
