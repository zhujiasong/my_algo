package leetcode

func ListPlusOne(head *ListNode) *ListNode {
	needPlusOne := _listPlusOne(head)
	if needPlusOne {
		return &ListNode{
			Val:  1,
			Next: head,
		}
	}
	return head
}

func _listPlusOne(head *ListNode) bool {
	if head == nil {
		return false
	}

	needPlusOne := _listPlusOne(head.Next)
	if head.Next == nil || needPlusOne {
		head.Val += 1
		if head.Val == 10 {
			head.Val = 0
			return true
		} else {
			return false
		}
	}

	return needPlusOne
}

func ListPlusOneV2(head *ListNode) *ListNode {
	head = ListTrverse(head)
	if head == nil {
		return head
	}

	h := head
	for h != nil {
		h.Val += 1
		if h.Val < 10 {
			break
		} else {
			h.Val = 0
		}
		h = h.Next
	}

	head = ListTrverse(head)
	if head.Val != 0 {
		return head
	}

	return &ListNode{
		Val:  1,
		Next: head,
	}
}
