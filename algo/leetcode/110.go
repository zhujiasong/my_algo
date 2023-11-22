package leetcode

import (
	"math"
	tree "my_algo/ds/binary_tree"
)

var isBalancedTree bool = true

func IsBalanced(root *tree.TreeNode) bool {
	if root == nil {
		return true
	}
	maxDepth(root)
	return isBalancedTree
}

func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)

	if int(math.Abs(float64(ld-rd))) > 1 {
		isBalancedTree = false
	}

	return max(ld, rd) + 1
}
