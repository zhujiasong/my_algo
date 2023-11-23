/*
129. Sum Root to Leaf Numbers
Medium
7.3K
116
Companies
You are given the root of a binary tree containing digits from 0 to 9 only.

Each root-to-leaf path in the tree represents a number.

For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
Return the total sum of all root-to-leaf numbers. Test cases are generated so that the answer will fit in a 32-bit integer.

A leaf node is a node with no children.
*/
package leetcode

import (
	tree "my_algo/ds/binary_tree"
)

var root2LeafSumTotal int
var root2LeafSum int

func RootToLeafSumNumbers(root *tree.TreeNode) int {
	root2LeafSumTotal, root2LeafSum = 0, 0
	rootToleafTraverse(root)
	return root2LeafSumTotal
}

func rootToleafTraverse(root *tree.TreeNode) {
	if root == nil {
		return
	}

	root2LeafSum = root2LeafSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		root2LeafSumTotal += root2LeafSum
	}

	rootToleafTraverse(root.Left)
	rootToleafTraverse(root.Right)
	root2LeafSum = (root2LeafSum - root.Val) / 10
}
