package main

import (
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	expect("#1", hasCycleFromSlices([]int{1, 2, 3, 4}, 1), true)
	expect("#2", hasCycleFromSlices([]int{1, 2}, -1), false)
	expect("#2", hasCycleFromSlices([]int{1, 2, 1}, -1), false)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func sliceToLinkedList(nums []int, index int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	var cycle *ListNode = nil
	if index == 0 {
		cycle = head
	}
	current := head
	for i := 1; i < len(nums); i++ {
		x := &ListNode{nums[i], nil}
		current.Next = x
		current = x
		if i == index {
			cycle = x
		}
	}
	if cycle != nil {
		current.Next = current
	}
	return head
}

func hasCycleFromSlices(nums []int, index int) bool {
	head := sliceToLinkedList(nums, index)
	return hasCycle(head)
}

// func hasCycle(head *ListNode) bool {
// 	m := make(map[*ListNode]bool)
// 	current := head
// 	for current != nil {
// 		if m[current] {
// 			return true
// 		}
// 		m[current] = true
// 		current = current.Next
// 	}
// 	return false
// }

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
