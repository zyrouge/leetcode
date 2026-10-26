package main

import (
	"fmt"
	"math"
	"os"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const TNNone = math.MaxInt

func main() {
	expect("#1", isValidBSTFromSlices([]int{2, 1, 3}), true)
	expect("#2", isValidBSTFromSlices([]int{1, 2, 3}), false)
	expect("#3", isValidBSTFromSlices([]int{2, 2, 2}), false)
	expect("#4", isValidBSTFromSlices([]int{0, -1000, 1000, TNNone, TNNone, 0}), false)
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

func isValidBSTFromSlices(nums []int) bool {
	root := sliceToTreeNode(nums)
	return isValidBST(root)
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTNode(root, math.MinInt, math.MaxInt)
}

func isValidBSTNode(node *TreeNode, left int, right int) bool {
	if node == nil {
		return true
	}
	if node.Val <= left || node.Val >= right {
		return false
	}
	return isValidBSTNode(node.Left, left, node.Val) && isValidBSTNode(node.Right, node.Val, right)
}
