/*
105. Construct Binary Tree from Preorder and Inorder Traversal
Medium
14.2K
438
Companies
Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
*/

package leetcode

import (
	tree "my_algo/ds/binary_tree"
)

var _buildTreeIndex int

// TODE, not compelete yet
func BuildTree(preorder []int, inorder []int) *tree.TreeNode {
	_buildTreeIndex = 0
	return _buildTree(preorder, inorder)
}

func _buildTree(preorder []int, inorder []int) *tree.TreeNode {
	if _buildTreeIndex == len(preorder) {
		return nil
	}

	val := preorder[_buildTreeIndex]
	_buildTreeIndex++

	root := new(tree.TreeNode)
	root.Val = val

	if len(inorder) <= 1 {
		return root
	}

	idx := getIndex(inorder, val)
	root.Left = _buildTree(preorder, inorder[:idx])
	root.Right = _buildTree(preorder, inorder[idx+1:])

	return root
}
