/*
19. Remove Nth Node From End of List
*/

package leetcode

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	fast, slow := head, head
	pre := head

	for ; n > 0; n-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
	}

	// to be deleted node is head
	if slow == head {
		return head.Next
	}

	// to be deleted node is tail
	if slow.Next == nil {
		pre.Next = nil
		return head
	}

	pre.Next = slow.Next
	return head
}
