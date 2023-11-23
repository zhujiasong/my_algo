/*
112. Path Sum
Easy
9.2K
1K
Companies
Given the root of a binary tree and an integer targetSum, return true if the
tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.
*/
package leetcode

import tree "my_algo/ds/binary_tree"

func HasPathSum(root *tree.TreeNode, targetSum int) bool {
	flg := false
	sum := 0
	order(root, targetSum, &sum, &flg)
	return flg
}

func order(root *tree.TreeNode, targetSum int, sum *int, flg *bool) {
	if root == nil {
		return
	}

	*sum += root.Val
	if root.Left == nil && root.Right == nil && *sum == targetSum {
		*flg = true
		return
	}

	order(root.Left, targetSum, sum, flg)
	order(root.Right, targetSum, sum, flg)
	*sum -= root.Val
}

func HasPathSumV2(root *tree.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	return HasPathSum(root.Left, targetSum-root.Val) ||
		HasPathSum(root.Right, targetSum-root.Val)
}
