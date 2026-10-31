package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	expectFunc("#1", cloneGraphFromSlices([][]int{{2}, {1, 3}, {2}}), [][]int{{2}, {1, 3}, {2}}, nestedSliceEqualFunc)
	expectFunc("#2", cloneGraphFromSlices([][]int{{}}), [][]int{{}}, nestedSliceEqualFunc)
	expectFunc("#2", cloneGraphFromSlices([][]int{}), [][]int{}, nestedSliceEqualFunc)
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

func adjacentSliceToNode(nums [][]int) *Node {
	if len(nums) == 0 {
		return nil
	}
	nodes := make(map[int]*Node)
	for i, adjacents := range nums {
		xv := i + 1
		x := nodes[xv]
		if x == nil {
			x = &Node{xv, make([]*Node, 0)}
			nodes[xv] = x
		}
		for _, yv := range adjacents {
			y := nodes[yv]
			if y == nil {
				y = &Node{yv, make([]*Node, 0)}
				nodes[yv] = y
			}
			x.Neighbors = append(x.Neighbors, y)
		}
		nodes[xv] = x
	}
	return nodes[1]
}

func nodeToAdjacentSlice(node *Node) [][]int {
	nums := make([][]int, 0)
	if node == nil {
		return nums
	}
	nodes := make(map[int]*Node)
	queue := []*Node{node}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		if visited[x.Val] {
			continue
		}
		visited[x.Val] = true
		nodes[x.Val] = x
		queue = append(queue, x.Neighbors...)
	}
	for i := range len(nodes) {
		xv := i + 1
		x := make([]int, len(nodes[xv].Neighbors))
		for j, y := range nodes[xv].Neighbors {
			x[j] = y.Val
		}
		nums = append(nums, x)
	}
	return nums
}

func cloneGraphFromSlices(nums [][]int) [][]int {
	node := adjacentSliceToNode(nums)
	cloned := cloneGraph(node)
	return nodeToAdjacentSlice(cloned)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	nodes := make(map[int]*Node)
	queue := []*Node{node}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		if visited[x.Val] {
			continue
		}
		visited[x.Val] = true
		xc := nodes[x.Val]
		if xc == nil {
			xc = &Node{x.Val, make([]*Node, 0)}
			nodes[x.Val] = xc
		}
		for _, y := range x.Neighbors {
			yc := nodes[y.Val]
			if yc == nil {
				yc = &Node{y.Val, make([]*Node, 0)}
				nodes[y.Val] = yc
			}
			xc.Neighbors = append(xc.Neighbors, yc)
		}
		queue = append(queue, x.Neighbors...)
	}
	return nodes[1]
}
