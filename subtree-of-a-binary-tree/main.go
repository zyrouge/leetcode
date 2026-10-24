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
	expect("#1", isSubtreeFromSlices([]int{1, 2, 3, 4, 5}, []int{2, 4, 5}), true)
	expect("#2", isSubtreeFromSlices([]int{1, 2, 3, 4, 5, TNNone, TNNone, 6}, []int{2, 4, 5}), false)
	expect("#3", isSubtreeFromSlices([]int{1, 1}, []int{1}), true)
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

func isSubtreeFromSlices(pNums []int, qNums []int) bool {
	root := sliceToTreeNode(pNums)
	subRoot := sliceToTreeNode(qNums)
	return isSubtree(root, subRoot)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}
