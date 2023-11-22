package leetcode

import tree "my_algo/ds/binary_tree"

func MaxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := MaxDepth(root.Left)
	maxRight := MaxDepth(root.Right)

	return max(maxLeft, maxRight) + 1
}
