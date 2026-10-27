package main

import (
	"fmt"
	"os"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const TNNone = -1

func main() {
	expect("#1", isSameTreeFromSlices([]int{2, 1, 3}, 1), 1)
	expect("#2", isSameTreeFromSlices([]int{4, 3, 5, 2, TNNone}, 4), 5)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func sliceToTreeNode(nums []int) *TreeNode {
	c := len(nums)
	if c == 0 {
		return nil
	}
	root := &TreeNode{nums[0], nil, nil}
	queue := []*TreeNode{root}
	i := 1
	for i < c && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if nums[i] != TNNone {
			node.Left = &TreeNode{nums[i], nil, nil}
			queue = append(queue, node.Left)
		}
		i++
		if i >= c {
			break
		}
		if nums[i] != TNNone {
			node.Right = &TreeNode{nums[i], nil, nil}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func isSameTreeFromSlices(nums []int, k int) int {
	root := sliceToTreeNode(nums)
	return kthSmallest(root, k)
}

// func kthSmallest(root *TreeNode, k int) int {
// 	nums := make([]int, 0)
// 	queue := []*TreeNode{root}
// 	for len(queue) > 0 {
// 		x := queue[0]
// 		queue = queue[1:]
// 		nums = append(nums, x.Val)
// 		if x.Left != nil {
// 			queue = append(queue, x.Left)
// 		}
// 		if x.Right != nil {
// 			queue = append(queue, x.Right)
// 		}
// 	}
// 	slices.Sort(nums)
// 	return nums[k-1]
// }

func kthSmallest(root *TreeNode, k int) int {
	n, result := -1, -1
	var dfs func(node *TreeNode, k int)
	dfs = func(node *TreeNode, k int) {
		if node.Left != nil {
			dfs(node.Left, k)
		}
		n++
		if n+1 > k {
			return
		}
		if n+1 == k {
			result = node.Val
			return
		}
		if node.Right != nil {
			dfs(node.Right, k)
		}
	}
	dfs(root, k)
	return result
}
