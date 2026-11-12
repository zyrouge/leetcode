package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	current1, current2 := l1, l2
	dummy := &ListNode{0, nil}
	parent := dummy
	for current1 != nil || current2 != nil || carry > 0 {
		if current1 != nil {
			carry += current1.Val
			current1 = current1.Next
		}
		if current2 != nil {
			carry += current2.Val
			current2 = current2.Next
		}
		parent.Next = &ListNode{carry % 10, nil}
		parent = parent.Next
		carry /= 10
	}
	return dummy.Next
}
