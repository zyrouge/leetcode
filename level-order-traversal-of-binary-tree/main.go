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

const TNNone = -101

func main() {
	expectFunc("#1", levelOrderFromSlices([]int{1, 2, 3, 4, 5, 6, 7}), [][]int{{1}, {2, 3}, {4, 5, 6, 7}}, nestedSliceEqualFunc)
	expectFunc("#2", levelOrderFromSlices([]int{1}), [][]int{{1}}, nestedSliceEqualFunc)
	expectFunc("#2", levelOrderFromSlices([]int{}), [][]int{}, nestedSliceEqualFunc)
}

func nestedSliceEqualFunc[T comparable](a, b [][]T) bool {
	return slices.EqualFunc(a, b, slices.Equal)
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

func levelOrderFromSlices(nums []int) [][]int {
	root := sliceToTreeNode(nums)
	return levelOrder(root)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		result = append(result, make([]int, 0))
		for range len(queue) {
			node := queue[0]
			queue = queue[1:]
			result[level] = append(result[level], node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
	}
	return result
}
