package list

import "fmt"

type sListNode struct {
	Val  int
	Next *sListNode
}

type sList struct {
	head, tail *sListNode
	len        int
}

func CreateSList(nums []int) *sList {
	l := new(sList)
	l.head = new(sListNode)

	if len(nums) == 0 {
		return l
	}

	l.len = len(nums)
	l.tail = l.head

	for _, num := range nums {
		node := &sListNode{
			Val: num,
		}

		l.tail.Next = node
		l.tail = node
	}

	return l
}

func (l *sList) Print() {
	if l == nil || l.head == nil {
		return
	}

	node := l.head.Next
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Printf("\n")
}

func (l *sList) IsSymmetryV2() bool {
	flg := true
	_isSymmetryV2(l.head.Next, &l.head.Next, &flg)
	return flg
}

func _isSymmetryV2(head *sListNode, node **sListNode, flg *bool) {
	if head == nil {
		return
	}

	_isSymmetryV2(head.Next, node, flg)
	if head.Val != (*node).Val {
		*flg = false
		return
	}
	*node = (*node).Next
}

func (l *sList) IsSymmetry() bool {
	if l.len <= 1 {
		return true
	}

	dummy := new(sListNode)
	slow, fast := l.head.Next, l.head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		tmp := slow
		slow = slow.Next

		tmp.Next = dummy.Next
		dummy.Next = tmp
	}

	var right *sListNode
	if fast == nil { // 节点数量为双数
		right = slow
	} else if fast.Next == nil { // 节点数量为单数
		right = slow.Next
	}

	return equal(dummy.Next, right)
}

func equal(l1, l2 *sListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1, l2 = l1.Next, l2.Next
	}

	return l1 == nil && l2 == nil
}

func (l *sList) Reverse() {
	node := reverseSlist(l.head.Next)
	l.head.Next = node
}

func reverseSlist(head *sListNode) *sListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseSlist(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
}
