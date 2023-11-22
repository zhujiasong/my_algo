package leetcode

import tree "my_algo/ds/binary_tree"

func IsSymmetric(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	return f(root.Left, root.Right)
}

func f(left, right *tree.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return f(left.Left, right.Right) && f(left.Right, right.Left)
}
