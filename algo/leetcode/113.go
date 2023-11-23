package leetcode

import tree "my_algo/ds/binary_tree"

func PathSum(root *tree.TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	trackArr := make([]int, 0)

	var backtrack func(*tree.TreeNode, int)
	backtrack = func(root *tree.TreeNode, targetSum int) {
		if root == nil {
			return
		}

		trackArr = append(trackArr, root.Val)
		if root.Left == nil && root.Right == nil && root.Val == targetSum {
			res = append(res, append([]int{}, trackArr...))
		}

		backtrack(root.Left, targetSum-root.Val)
		backtrack(root.Right, targetSum-root.Val)

		trackArr = trackArr[:len(trackArr)-1]
	}

	backtrack(root, targetSum)
	return res
}
