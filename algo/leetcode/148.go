/*
148. sort list
*/

package leetcode

// not finished, todo
func sortList(head *ListNode) *ListNode {
	return _sortList(head, nil)
}

func _sortList(head *ListNode, end *ListNode) *ListNode {
	if head == end || head.Next == end {
		return head
	}

	newHead := partitionListNode(head, nil)
	_sortList(newHead, head)
	_sortList(head.Next, nil)

	return newHead
}

func partitionListNode(head *ListNode, end *ListNode) *ListNode {
	leftTail, dummyRight := new(ListNode), new(ListNode)
	dummyLeft := leftTail
	curr := head

	for curr != end {
		if curr.Val < head.Val {
			leftTail.Next = curr
			leftTail = leftTail.Next
		} else {
			curr.Next = dummyRight.Next
			dummyRight.Next = curr
		}
		curr = curr.Next
	}

	leftTail.Next = head
	head.Next = dummyRight.Next

	return dummyLeft.Next
}
