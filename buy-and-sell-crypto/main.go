package main

import "fmt"

func main() {
	// prices := []int{10, 1, 5, 6, 7, 1} // 6
	prices := []int{10, 8, 7, 5, 2} // 0
	fmt.Printf("%d\n", maxProfit(prices))
}

func maxProfit(prices []int) int {
	i, j, result := 0, 1, 0
	for j < len(prices) {
		x, y := prices[i], prices[j]
		profit := y - x
		if profit > result {
			result = profit
		}
		if y < x {
			i = j
			continue
		}
		j++
	}
	return result
}
