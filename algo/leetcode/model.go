package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums []int) *ListNode {
	dummy := new(ListNode)
	h := dummy

	for _, num := range nums {
		node := &ListNode{
			Val: num,
		}
		h.Next = node
		h = h.Next
	}

	return dummy.Next
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}

func ListTrverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	h := dummy
	for head != nil {
		t := head
		head = head.Next
		t.Next = h.Next
		h.Next = t
	}
	return dummy.Next
}
