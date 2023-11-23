/*
543. Diameter of Binary Tree
Easy
Topics
Companies
Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.



Example 1:


Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
Example 2:

Input: root = [1,2]
Output: 1
*/

package leetcode

import tree "my_algo/ds/binary_tree"

func DiameterOfBinaryTree(root *tree.TreeNode) int {
	var maxDiameter int
	var dfs func(*tree.TreeNode) int

	dfs = func(root *tree.TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			return 1
		}

		l := dfs(root.Left)
		r := dfs(root.Right)

		maxDiameter = max(maxDiameter, l+r)

		return max(l, r) + 1
	}

	dfs(root)

	return maxDiameter
}
