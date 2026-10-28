package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const TNNone = -1001

func main() {
	expectFunc("#1", invertTreeFromSlices([]int{1, 2, 3, TNNone, TNNone, 4, 5}), []int{1, 2, 3, TNNone, TNNone, 4, 5}, slices.Equal)
	expectFunc("#2", invertTreeFromSlices([]int{}), []int{}, slices.Equal)
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

func invertTreeFromSlices(nums []int) []int {
	root := sliceToTreeNode(nums)
	codec := Constructor()
	serialized := codec.serialize(root)
	deserialized := codec.deserialize(serialized)
	return treeNodeToSlice(deserialized)
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var result strings.Builder
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		if x == nil {
			result.WriteString(",nil")
			continue
		}
		result.WriteByte(',')
		result.WriteString(strconv.Itoa(x.Val))
		queue = append(queue, x.Left, x.Right)
	}
	return result.String()[1:]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	var readValue func() int
	i := 0
	readValue = func() int {
		var b strings.Builder
		for i < len(data) && data[i] != ',' {
			b.WriteByte(data[i])
			i++
		}
		i++
		s := b.String()
		if s == "nil" {
			return TNNone
		}
		v, _ := strconv.Atoi(s)
		return v
	}
	root := &TreeNode{readValue(), nil, nil}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		x := queue[0]
		queue = queue[1:]
		lv, rv := readValue(), readValue()
		if lv != TNNone {
			x.Left = &TreeNode{lv, nil, nil}
			queue = append(queue, x.Left)
		}
		if rv != TNNone {
			x.Right = &TreeNode{rv, nil, nil}
			queue = append(queue, x.Right)
		}
	}
	return root
}
