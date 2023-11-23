/*
148. Sort List
Solved
Medium
Topics
Companies
Given the head of a linked list, return the list after sorting it in ascending order.
*/

package leetcode

// TODO,not complete
func mergesortList(head *ListNode) *ListNode {
	return _mergeSortList(head, nil)
}

func _mergeSortList(start, end *ListNode) *ListNode {
	if start == end || start.Next == end {
		return start
	}

	slow, fast := start, start
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	_mergeSortList(start, slow)
	_mergeSortList(slow.Next, end)
	return _mergeList(start, slow, slow.Next, end)
}

func _mergeList(start1, end1, start2, end2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy

	for start1 != end1 && start2 != end2 {
		var tmp *ListNode
		if start1.Val < start2.Val {
			tmp = start1
			start1 = start1.Next
		} else {
			tmp = start2
			start2 = start2.Next
		}
		curr.Next = tmp
		curr = tmp
	}
	curr.Next = nil

	if start1 != end1 {
		curr.Next = start1
	}
	if start2 != end2 {
		curr.Next = start2
	}

	return dummy.Next
}
