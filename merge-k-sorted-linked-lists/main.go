package main

import (
	"fmt"
	"math"
	"os"
	"slices"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	expectFunc("#1", mergeKListsFromSlices([][]int{{1, 2, 4}, {1, 3, 5}, {3, 6}}), []int{1, 1, 2, 3, 3, 4, 5, 6}, slices.Equal)
	expectFunc("#2", mergeKListsFromSlices([][]int{}), []int{}, slices.Equal)
	expectFunc("#3", mergeKListsFromSlices([][]int{{}}), []int{}, slices.Equal)
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

func mergeKListsFromSlices(nums [][]int) []int {
	lists := make([]*ListNode, len(nums))
	for i, x := range nums {
		lists[i] = sliceToLinkedList(x)
	}
	result := mergeKLists(lists)
	return linkedListToSlice(result)
}

// type MinHeap []*ListNode

// func (h MinHeap) Len() int {
// 	return len(h)
// }

// func (h MinHeap) Less(i, j int) bool {
// 	return h[i].Val < h[j].Val
// }

// func (h MinHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func (h *MinHeap) Pop() any {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[:n-1]
// 	return x
// }

// func (h *MinHeap) Push(x any) {
// 	*h = append(*h, x.(*ListNode))
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	pq := &MinHeap{}
// 	heap.Init(pq)
// 	for _, x := range lists {
// 		if x != nil {
// 			heap.Push(pq, x)
// 		}
// 	}
// 	if pq.Len() == 0 {
// 		return nil
// 	}
// 	head := heap.Pop(pq).(*ListNode)
// 	if head.Next != nil {
// 		heap.Push(pq, head.Next)
// 	}
// 	head.Next = nil
// 	parent := head
// 	for pq.Len() > 0 {
// 		current := heap.Pop(pq).(*ListNode)
// 		if current.Next != nil {
// 			heap.Push(pq, current.Next)
// 		}
// 		current.Next = nil
// 		parent.Next = current
// 		parent = parent.Next
// 	}
// 	return head
// }

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	return mergeKListsDivide(lists, 0, n-1)
}

func mergeKListsDivide(lists []*ListNode, start, end int) *ListNode {
	if start > end {
		return nil
	}
	if start == end {
		return lists[start]
	}
	mid := start + ((end - start) / 2)
	list1 := mergeKListsDivide(lists, start, mid)
	list2 := mergeKListsDivide(lists, mid+1, end)
	return mergeLists(list1, list2)
}

func mergeLists(list1, list2 *ListNode) *ListNode {
	dummy := &ListNode{math.MinInt, nil}
	parent := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			parent.Next = list1
			list1 = list1.Next
		} else {
			parent.Next = list2
			list2 = list2.Next
		}
		parent = parent.Next
	}
	if list1 != nil {
		parent.Next = list1
	}
	if list2 != nil {
		parent.Next = list2
	}
	return dummy.Next
}
