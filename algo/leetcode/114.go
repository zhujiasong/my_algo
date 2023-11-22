/*
114. Flatten Binary Tree to Linked List
Medium
11.6K
541
Companies
Given the root of a binary tree, flatten the tree into a "linked list":

The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
The "linked list" should be in the same order as a pre-order traversal of the binary tree.
*/
package leetcode

import (
	tree "my_algo/ds/binary_tree"
)

func Flatten(root *tree.TreeNode) {
	if root == nil {
		return
	}

	Flatten(root.Left)
	Flatten(root.Right)

	right := root.Right
	root.Right = root.Left
	root.Left = nil

	r := root
	for r != nil && r.Right != nil {
		r = r.Right
	}
	r.Right = right
}
