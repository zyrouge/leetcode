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
	expect("#1", lowestCommonAncestorFromSlices([]int{5, 3, 8, 1, 4, 7, 9, TNNone, 2}, 3, 8), 5)
	expect("#2", lowestCommonAncestorFromSlices([]int{5, 3, 8, 1, 4, 7, 9, TNNone, 2}, 3, 4), 3)
	expect("#2", lowestCommonAncestorFromSlices([]int{5, 3, 8, 1, 4, 7, 9, TNNone, 2}, 1, 2), 1)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func sliceToTreeNodes(nums []int, p, q int) (*TreeNode, *TreeNode, *TreeNode) {
	c := len(nums)
	if c == 0 {
		return nil, nil, nil
	}
	root := &TreeNode{nums[0], nil, nil}
	var pRoot, qRoot *TreeNode = nil, nil
	updatePQRoot := func(node *TreeNode) {
		if node.Val == p {
			pRoot = node
		}
		if node.Val == q {
			qRoot = node
		}
	}
	queue := []*TreeNode{root}
	i := 1
	for i < c && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		updatePQRoot(node)
		if nums[i] != TNNone {
			node.Left = &TreeNode{nums[i], nil, nil}
			queue = append(queue, node.Left)
			updatePQRoot(node.Left)
		}
		i++
		if i >= c {
			break
		}
		if nums[i] != TNNone {
			node.Right = &TreeNode{nums[i], nil, nil}
			queue = append(queue, node.Right)
			updatePQRoot(node.Right)
		}
		i++
	}
	return root, pRoot, qRoot
}

func lowestCommonAncestorFromSlices(nums []int, p int, q int) int {
	root, pRoot, qRoot := sliceToTreeNodes(nums, p, q)
	result := lowestCommonAncestor(root, pRoot, qRoot)
	if result == nil {
		return TNNone
	}
	return result.Val
}

// type TreeNodePair struct {
// 	Parent *TreeNodePair
// 	Node   *TreeNode
// 	Depth  int
// }

// func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
// 	pairs := make(map[int]*TreeNodePair)
// 	stack := []*TreeNodePair{{nil, root, 0}}
// 	for len(stack) > 0 {
// 		pair := stack[0]
// 		stack = stack[1:]
// 		pairs[pair.Node.Val] = pair
// 		if pairs[p.Val] != nil && pairs[q.Val] != nil {
// 			break
// 		}
// 		if pair.Node.Left != nil {
// 			stack = append(stack, &TreeNodePair{pair, pair.Node.Left, pair.Depth + 1})
// 		}
// 		if pair.Node.Right != nil {
// 			stack = append(stack, &TreeNodePair{pair, pair.Node.Right, pair.Depth + 1})
// 		}
// 	}
// 	pPair, qPair := pairs[p.Val], pairs[q.Val]
// 	for pPair.Depth > qPair.Depth {
// 		pPair = pPair.Parent
// 	}
// 	for qPair.Depth > pPair.Depth {
// 		qPair = qPair.Parent
// 	}
// 	// fmt.Printf("%d:%d %d:%d\n", pPair.Node.Val, pPair.Depth, qPair.Node.Val, qPair.Depth)
// 	for pPair.Node.Val != qPair.Node.Val {
// 		pPair = pPair.Parent
// 		qPair = qPair.Parent
// 	}
// 	return pPair.Node
// }

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	parent := root
	for {
		// fmt.Println(parent)
		if p.Val < parent.Val && q.Val < parent.Val {
			parent = parent.Left
			continue
		}
		if p.Val > parent.Val && q.Val > parent.Val {
			parent = parent.Right
			continue
		}
		break
	}
	return parent
}
