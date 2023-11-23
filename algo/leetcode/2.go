/*
2. Add Two Numbers
Attempted
Medium
Topics
Companies
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/

package leetcode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	dummy := new(ListNode)
	h := dummy
	var val, plus int
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + plus
		if sum < 10 {
			val, plus = sum, 0
		} else {
			val, plus = sum-10, 1
		}
		node := &ListNode{
			Val: val,
		}
		h.Next = node
		h = h.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	lists := []*ListNode{l1, l2}
	for _, l := range lists {
		for l != nil {
			val = l.Val + plus
			if val < 10 {
				plus = 0
			} else {
				plus = 1
				val = 0
			}
			node := &ListNode{
				Val: val,
			}
			h.Next = node
			h = h.Next

			l = l.Next
		}
	}

	if plus != 0 {
		h.Next = &ListNode{
			Val: 1,
		}
	}
	return dummy.Next
}
