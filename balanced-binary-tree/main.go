package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		left, leftb := dfs(node.Left)
		right, rightb := dfs(node.Right)
		if !leftb || !rightb || abs(left-right) > 1 {
			return 0, false
		}
		return 1 + max(left, right), true
	}
	_, result := dfs(root)
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
