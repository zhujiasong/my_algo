/*
160. Intersection of Two Linked Lists
Solved
Easy
Topics
Companies
Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.
*/

package leetcode

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	h1, h2 := headA, headB
	for h1 != h2 {
		if h1 == nil {
			h1 = headB
		} else {
			h1 = h1.Next
		}

		if h2 == nil {
			h2 = headA
		} else {
			h2 = h2.Next
		}
	}

	return h1
}
