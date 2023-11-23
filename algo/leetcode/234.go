/*
234. Palindrome Linked List
Easy
Topics
Companies
Given the head of a singly linked list, return true if it is a
palindrome
 or false otherwise.



Example 1:


Input: head = [1,2,2,1]
Output: true
Example 2:


Input: head = [1,2]
Output: false


Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9
*/

package leetcode

func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	var slow, fast, dummy *ListNode
	dummy = new(ListNode)
	slow, fast = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tmp := slow
		slow = slow.Next

		tmp.Next = dummy.Next
		dummy.Next = tmp
	}

	left, right := dummy.Next, slow.Next
	// 链表节点个数是双数
	if fast == nil {
		right = slow
	}

	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next
	}

	return left == nil && right == nil
}

func IsPalindromeV2(head *ListNode) bool {
	isPalindromeFlg := true
	_ListTraverse(head, &head, &isPalindromeFlg)
	return isPalindromeFlg
}

func _ListTraverse(head *ListNode, node **ListNode, isPalindromeFlg *bool) {
	if head == nil {
		return
	}

	_ListTraverse(head.Next, node, isPalindromeFlg)
	if head.Val != (*node).Val {
		*isPalindromeFlg = false
	}
	(*node) = (*node).Next
}
