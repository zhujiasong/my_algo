/*
116. Populating Next Right Pointers in Each Node
Medium
9.3K
286
Companies
You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.
*/

package leetcode

type Node struct {
	Val               int
	Left, Right, Next *Node
}

func ConnectBTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	size := 1

	for len(queue) != 0 {
		tmpNode := queue[0]

		if tmpNode.Left != nil {
			queue = append(queue, tmpNode.Left)
		}
		if tmpNode.Right != nil {
			queue = append(queue, tmpNode.Right)
		}

		queue = queue[1:]
		size--
		if size == 0 {
			tmpNode.Next = nil
			size = len(queue)
		} else {
			tmpNode.Next = queue[0]
		}
	}

	return root
}
