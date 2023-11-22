package leetcode

import (
	tree "my_algo/ds/binary_tree"
	"my_algo/ds/list"
)

func SortedListToBST(head *list.ListNode) *tree.TreeNode {
	if head == nil {
		return nil
	}

	midPrev, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		midPrev = slow
		slow = slow.Next
	}
	midPrev.Next = nil

	midNode := slow
	left, right := head, slow.Next
	// no left node
	if midNode == midPrev {
		left = nil
	}

	root := new(tree.TreeNode)
	root.Val = midNode.Elem.(int)
	root.Left = SortedListToBST(left)
	root.Right = SortedListToBST(right)

	return root
}
