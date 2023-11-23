/*
404. Sum of Left Leaves
Easy
4.9K
282
Companies
Given the root of a binary tree, return the sum of all left leaves.

A leaf is a node with no children. A left leaf is a leaf that is the left child of another node.

Input: root = [3,9,20,null,null,15,7]
Output: 24
Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.
*/

package leetcode

import tree "my_algo/ds/binary_tree"

var _sumOfLeftLeavesNum int

func SumOfLeftLeaves(root *tree.TreeNode) int {
	_sumOfLeftLeavesNum = 0
	sumOfLeftLeavesTraverse(root)
	return _sumOfLeftLeavesNum
}

func sumOfLeftLeavesTraverse(root *tree.TreeNode) {
	if root == nil {
		return
	}

	sumOfLeftLeavesTraverse(root.Left)
	sumOfLeftLeavesTraverse(root.Right)

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		_sumOfLeftLeavesNum += root.Left.Val
	}
}

func SumOfLeftLeavesRecursion(root *tree.TreeNode) int {
	return _sumOfLeftLeavesRecursion(root, false)
}

func _sumOfLeftLeavesRecursion(root *tree.TreeNode, isLeftNode bool) int {
	if root == nil {
		return 0
	}
	if isLeftNode && root.Left == nil && root.Right == nil {
		return root.Val
	}

	left := _sumOfLeftLeavesRecursion(root.Left, true)
	right := _sumOfLeftLeavesRecursion(root.Right, false)

	return left + right
}

func SumOfLeftLeavesRecursionV(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	l := 0
	lnode := root.Left
	if lnode != nil && lnode.Left == nil && lnode.Right == nil {
		l = lnode.Val
	}

	left := SumOfLeftLeavesRecursionV(root.Left)
	right := SumOfLeftLeavesRecursionV(root.Right)
	return l + left + right
}
