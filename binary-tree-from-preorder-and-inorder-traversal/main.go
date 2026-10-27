package main

import (
	"fmt"
	"math"
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
	expectFunc("#3", buildTreeToSlices([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}), []int{3, 9, 20, TNNone, TNNone, 15, 7}, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func treeNodeToSlice(root *TreeNode) []int {
	result := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, TNNone)
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(result) > 0 && result[len(result)-1] == TNNone {
		result = result[:len(result)-1]
	}
	return result
}

func buildTreeToSlices(preorder, inorder []int) []int {
	result := buildTree(preorder, inorder)
	return treeNodeToSlice(result)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	pi, ii := 0, 0
	var dfs func(limit int) *TreeNode
	dfs = func(limit int) *TreeNode {
		if pi >= len(preorder) {
			return nil
		}
		if inorder[ii] == limit {
			ii++
			return nil
		}
		root := &TreeNode{preorder[pi], nil, nil}
		pi++
		// fmt.Println(root.Val, preorder[pi:], inorder[ii:], limit)
		root.Left = dfs(root.Val)
		root.Right = dfs(limit)
		return root
	}
	return dfs(math.MaxInt)
}
