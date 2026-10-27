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

const TNNone = -1001

func main() {
	expect("#1", maxPathSumFromSlice([]int{1, 2, 3}), 6)
	expect("#2", maxPathSumFromSlice([]int{-15, 10, 20, TNNone, TNNone, 15, 5, -5}), 40)
	expect("#3", maxPathSumFromSlice([]int{2, -1, -2}), 2)
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

func maxPathSumFromSlice(nums []int) int {
	root := sliceToTreeNode(nums)
	return maxPathSum(root)
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m, sm := maxPathSumSubtree(root)
	return max(m, sm)
}

func maxPathSumSubtree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	m := root.Val
	ls, rs := 0, 0
	if root.Left != nil {
		temp := 0
		temp, ls = maxPathSumSubtree(root.Left)
		m = max(m, temp)
	}
	if root.Right != nil {
		temp := 0
		temp, rs = maxPathSumSubtree(root.Right)
		m = max(m, temp)
	}
	m = max(m, root.Val+ls+rs)
	return m, max(root.Val, root.Val+max(ls, rs))
}

// func maxPathSum(root *TreeNode) int {
// 	result := 0
// 	var dfs func(root *TreeNode) int
// 	dfs = func(root *TreeNode) int {
// 		if root == nil {
// 			return 0
// 		}
// 		ls, rs := 0, 0
// 		if root.Left != nil {
// 			ls = dfs(root.Left)
// 		}
// 		if root.Right != nil {
// 			rs = dfs(root.Right)
// 		}
// 		result = max(result, root.Val+ls+rs)
// 		return root.Val + max(ls, rs)
// 	}
// 	dfs(root)
// 	return result
// }
