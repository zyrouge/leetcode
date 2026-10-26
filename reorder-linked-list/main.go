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
	expectFunc("#1", reorderListFromSlice([]int{0, 1, 2, 3, 4, 5, 6}), []int{0, 6, 1, 5, 2, 4, 3}, slices.Equal)
	expectFunc("#2", reorderListFromSlice([]int{2, 4, 6, 8}), []int{2, 8, 4, 6}, slices.Equal)
	expectFunc("#3", reorderListFromSlice([]int{2, 4, 6, 8, 10}), []int{2, 10, 4, 8, 6}, slices.Equal)
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

func reorderListFromSlice(nums []int) []int {
	head := sliceToLinkedList(nums)
	reorderList(head)
	return linkedListToSlice(head)
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow := head // mid
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	previous, current := (*ListNode)(nil), slow.Next
	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}
	slow.Next = nil
	part1 := head
	part2 := previous
	for part1 != nil && part2 != nil {
		next1 := part1.Next
		next2 := part2.Next
		part1.Next = part2
		part2.Next = next1
		part1 = next1
		part2 = next2
	}
}
