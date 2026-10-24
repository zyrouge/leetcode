package main

import (
	"fmt"
	"os"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	expectFunc("#1", invertTreeFromSlices([]int{1, 2, 3, 4, 5, 6, 7}), []int{1, 3, 2, 7, 6, 5, 4}, slices.Equal)
	expectFunc("#2", invertTreeFromSlices([]int{3, 2, 1}), []int{3, 1, 2}, slices.Equal)
	expectFunc("#3", invertTreeFromSlices([]int{}), []int{}, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
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
		node.Left = &TreeNode{nums[i], nil, nil}
		queue = append(queue, node.Left)
		i++
		if i >= c {
			break
		}
		node.Right = &TreeNode{nums[i], nil, nil}
		queue = append(queue, node.Right)
		i++
	}
	return root
}

func treeNodeToSlice(root *TreeNode) []int {
	nums := make([]int, 0)
	if root == nil {
		return make([]int, 0)
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nums = append(nums, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return nums
}

func invertTreeFromSlices(nums []int) []int {
	root := sliceToTreeNode(nums)
	result := invertTree(root)
	return treeNodeToSlice(result)
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		temp := node.Left
		node.Left = node.Right
		node.Right = temp
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}
