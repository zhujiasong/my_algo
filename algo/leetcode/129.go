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

func SumNumbers(root *tree.TreeNode) int {
	var sum, totalSum int
	var backtrack func(*tree.TreeNode)
	backtrack = func(root *tree.TreeNode) {
		if root == nil {
			return
		}

		sum = 10*sum + root.Val
		if root.Left == nil && root.Right == nil {
			totalSum += sum
		}

		backtrack(root.Left)
		backtrack(root.Right)

		sum /= 10
	}

	backtrack(root)
	return totalSum
}
