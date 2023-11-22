/*
226. Invert Binary Tree
Easy
13.3K
191
Companies
Given the root of a binary tree, invert the tree, and return its root.
*/
package leetcode

import tree "my_algo/ds/binary_tree"

func InvertTree(root *tree.TreeNode) *tree.TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	left := InvertTree(root.Left)
	right := InvertTree(root.Right)

	root.Left = right
	root.Right = left

	return root
}
