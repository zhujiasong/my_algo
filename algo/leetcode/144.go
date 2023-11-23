/*
144. Binary Tree Preorder Traversal
*/

package leetcode

import (
	tree "my_algo/ds/binary_tree"
)

func PreorderTraversal(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	stack := []*tree.TreeNode{root}

	for len(stack) != 0 {
		t := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, t.Val)

		if t.Right != nil {
			stack = append(stack, t.Right)
		}
		if t.Left != nil {
			stack = append(stack, t.Left)
		}
	}
	return ret
}
