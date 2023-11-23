package list

import "fmt"

type ListNode struct {
	Elem       interface{}
	prev, Next *ListNode
}

// List represents double linked list
// head node is not real head node, it's just a dummy node
type List struct {
	head, tail *ListNode
	len        int
}

func New(elems []interface{}) *List {
	lst := new(List)
	lst.len = len(elems)
	lst.head = new(ListNode)

	if lst.len == 0 {
		return lst
	}

	lst.tail = lst.head
	for _, elem := range elems {
		t := new(ListNode)
		t.Elem = elem

		t.Next = lst.tail.Next
		lst.tail.Next = t
		t.prev = lst.tail
		lst.tail = t
	}

	return lst
}

func (l *List) IsEmpty() bool {
	return l.tail == nil
}

func (l *List) AddNodeHead(node *ListNode) {
	if l.IsEmpty() {
		l.tail = node
	} else {
		l.head.Next.prev = node
	}
	node.Next = l.head.Next
	l.head.Next = node
	node.prev = l.head

	l.len++
}

func (l *List) AddNodeTail(node *ListNode) {
	if l.IsEmpty() {
		l.AddNodeHead(node)
		return
	}

	l.tail.Next = node
	node.prev = l.tail
	l.tail = node

	l.len++
}

func (l *List) Front() *ListNode {
	return l.head.Next
}

func (l *List) Tail() *ListNode {
	return l.tail
}

func (l *List) Next(node *ListNode) *ListNode {
	return node.Next
}

func (l *List) Prev(node *ListNode) *ListNode {
	return node.prev
}

func (l *List) DeleteNode(node *ListNode) {
	if node == nil {
		return
	}

	if node == l.tail {
		l.tail = l.tail.prev
		l.tail.Next = nil

		// empty list
		if l.tail == l.head {
			l.tail = nil
		}
	} else {
		node.prev.Next = node.Next
		node.Next.prev = node.prev
	}

	l.len--
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Print() {
	for curr := l.head.Next; curr != nil; curr = curr.Next {
		fmt.Printf("%v ", curr.Elem)
	}
	fmt.Printf("\n")
}

func (l *List) PrintReverse() {
	for curr := l.tail; curr != nil && curr != l.head; curr = curr.prev {
		fmt.Printf("%v ", curr.Elem)
	}
	fmt.Printf("\n")
}
