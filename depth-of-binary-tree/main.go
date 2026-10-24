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

const TNNone = -101

func main() {
	expect("#1", maxDepthFromSlices([]int{1, 2, 3, TNNone, TNNone, 4}), 3)
	expect("#2", maxDepthFromSlices([]int{}), 0)
	expect("#3", maxDepthFromSlices([]int{1, 2, 3, 4, TNNone, TNNone, 5}), 3)
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

func maxDepthFromSlices(nums []int) int {
	root := sliceToTreeNode(nums)
	return maxDepth(root)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0
	queue := []*TreeNode{root}
	for {
		c := len(queue)
		if c == 0 {
			break
		}
		for range c {
			x := queue[0]
			queue = queue[1:]
			if x.Left != nil {
				queue = append(queue, x.Left)
			}
			if x.Right != nil {
				queue = append(queue, x.Right)
			}
		}
		result++
	}
	return result
}
