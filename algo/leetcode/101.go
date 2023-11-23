package leetcode

import tree "my_algo/ds/binary_tree"

func IsSymmetric(root *tree.TreeNode) bool {
	var dfs func(left, right *tree.TreeNode) bool
	dfs = func(left, right *tree.TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}

		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}

	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}
