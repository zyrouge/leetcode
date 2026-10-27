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

const TNNone = -1001

func main() {
	expectFunc("#1", buildTreeToSlices([]int{1, 2, 3, 4}, []int{2, 1, 3, 4}), []int{1, 2, 3, TNNone, TNNone, TNNone, 4}, slices.Equal)
	expectFunc("#2", buildTreeToSlices([]int{1}, []int{1}), []int{1}, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
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

func buildTreeToSlices(preorder, inorder []int) []int {
	result := buildTree(preorder, inorder)
	return treeNodeToSlice(result)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// for len(preorder) > 0 || len(inorder) > 0 {
	// in  - [2, 1, -, 3, -, 4, -]
	// pre - [1, 2, 3, 4]
	// }
	var build func() *TreeNode
	build = func() *TreeNode {

		return nil
	}
	return build()
}
