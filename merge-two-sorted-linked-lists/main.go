package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := sliceToLinkedList([]int{1, 2, 4})
	list2 := sliceToLinkedList([]int{1, 3, 5})
	result := mergeTwoLists(list1, list2)
	fmt.Println(linkedListToSlice(result))
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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	var head *ListNode
	if list2 != nil {
		if list1 != nil && list1.Val < list2.Val {
			head = list1
			list1 = list1.Next
		} else {
			head = list2
			list2 = list2.Next
		}
	} else {
		head = list1
		list1 = list1.Next
	}
	current := head
	for {
		if list1 != nil && (list2 == nil || list1.Val < list2.Val) {
			current.Next = list1
			current = current.Next
			list1 = list1.Next
			continue
		}
		if list2 != nil {
			current.Next = list2
			current = current.Next
			list2 = list2.Next
			continue
		}
		break
	}
	current.Next = nil
	return head
}
