package leetcode

import tree "my_algo/ds/binary_tree"

var targetSumpaths [][]int
var targetSumpath []int

func PathSum(root *tree.TreeNode, targetSum int) [][]int {
	targetSumpaths = make([][]int, 0)
	targetSumpath = make([]int, 0)
	targetPathSumTraverse(root, targetSum)
	return targetSumpaths
}

func targetPathSumTraverse(root *tree.TreeNode, targetSum int) {
	if root == nil {
		return
	}

	targetSumpath = append(targetSumpath, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		targetSumpaths = append(targetSumpaths, append([]int{}, targetSumpath...))
	}
	targetPathSumTraverse(root.Left, targetSum-root.Val)
	targetPathSumTraverse(root.Right, targetSum-root.Val)
	targetSumpath = targetSumpath[:len(targetSumpath)-1]
}
