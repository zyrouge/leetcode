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
	expect("#1", isSameTreeFromSlices([]int{1, 2, 3}, []int{1, 2, 3}), true)
	expect("#2", isSameTreeFromSlices([]int{4, 7}, []int{4, TNNone, 7}), false)
	expect("#3", isSameTreeFromSlices([]int{1}, []int{1, TNNone, 2}), false)
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

func isSameTreeFromSlices(pNums []int, qNums []int) bool {
	p := sliceToTreeNode(pNums)
	q := sliceToTreeNode(qNums)
	return isSameTree(p, q)
}

// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	pQueue, qQueue := []*TreeNode{p}, []*TreeNode{q}
// 	for len(pQueue) > 0 && len(qQueue) > 0 {
// 		x, y := pQueue[0], qQueue[0]
// 		pQueue, qQueue = pQueue[1:], qQueue[1:]
// 		if x == nil && y == nil {
// 			continue
// 		}
// 		if (x == nil && y != nil) || (y == nil && x != nil) || x.Val != y.Val {
// 			return false
// 		}
// 		pQueue = append(pQueue, x.Left, x.Right)
// 		qQueue = append(qQueue, y.Left, y.Right)
// 	}
// 	return true
// }

type TreeNodePair struct {
	First  *TreeNode
	Second *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNodePair{{p, q}}
	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]
		x, y := pair.First, pair.Second
		if x == nil && y == nil {
			continue
		}
		if (x == nil && y != nil) || (y == nil && x != nil) || x.Val != y.Val {
			return false
		}
		queue = append(queue, &TreeNodePair{x.Left, y.Left}, &TreeNodePair{x.Right, y.Right})
	}
	return true
}
