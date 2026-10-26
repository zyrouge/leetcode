package main

import (
	"fmt"
	"os"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	expectFunc("#1", removeNthFromEndFromSlice([]int{1, 2, 3, 4}, 2), []int{1, 2, 4}, slices.Equal)
	expectFunc("#2", removeNthFromEndFromSlice([]int{5}, 1), []int{}, slices.Equal)
	expectFunc("#3", removeNthFromEndFromSlice([]int{1, 2}, 2), []int{2}, slices.Equal)
	expectFunc("#4", removeNthFromEndFromSlice([]int{1, 2}, 1), []int{1}, slices.Equal)
}

func expectFunc[T any](message string, got T, expected T, equals func(T, T) bool) {
	if !equals(got, expected) {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func sliceToLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	current := head
	for i := 1; i < len(nums); i++ {
		x := &ListNode{nums[i], nil}
		current.Next = x
		current = current.Next
	}
	return head
}

func linkedListToSlice(head *ListNode) []int {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func removeNthFromEndFromSlice(nums []int, n int) []int {
	head := sliceToLinkedList(nums)
	result := removeNthFromEnd(head, n)
	return linkedListToSlice(result)
}

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
// 	size := 0
// 	current := head
// 	for current != nil {
// 		size++
// 		current = current.Next
// 	}
// 	if n == size {
// 		return head.Next
// 	}
// 	current = head
// 	for i := 1; i < size-n; i++ {
// 		current = current.Next
// 	}
// 	current.Next = current.Next.Next
// 	return head
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	for range n {
		first = first.Next
	}
	if first == nil {
		return head.Next
	}
	second := head
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return head
}
