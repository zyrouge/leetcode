package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	// nums := []int{0, -1, -1}
	// nums := []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	l := len(nums)
	for i := range l {
		for j := i + 1; j < l; j++ {
			x := nums[i]
			y := nums[j]
			if x > y {
				nums[i], nums[j] = y, x
			}
		}
	}
	result := make([][]int, 0)
	for i := range l {
		x := nums[i]
		if i > 0 && x == nums[i-1] {
			continue
		}
		j, k := i+1, l-1
		for j < k {
			y := nums[j]
			z := nums[k]
			sum := x + y + z
			if sum > 0 {
				k--
				continue
			}
			if sum < 0 {
				j++
				continue
			}
			result = append(result, []int{x, y, z})
			j++
			k--
			for j < k && nums[j-1] == nums[j] {
				j++
			}
		}
	}
	return result
}
