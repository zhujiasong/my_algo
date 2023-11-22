package leetcode

import (
	"fmt"
	tree "my_algo/ds/binary_tree"
)

var bTreePaths []string
var bTreePath []int

func BinaryTreePaths(root *tree.TreeNode) []string {
	bTreePaths = make([]string, 0)
	bTreePath = make([]int, 0)
	pathSumTraverse(root)
	return bTreePaths
}

func pathSumTraverse(root *tree.TreeNode) {
	if root == nil {
		return
	}

	bTreePath = append(bTreePath, root.Val)
	if root.Left == nil && root.Right == nil {
		var s string
		for _, v := range bTreePath {
			s = fmt.Sprintf("%s%d->", s, v)
		}
		s = s[:len(s)-2]
		bTreePaths = append(bTreePaths, s)
	}
	pathSumTraverse(root.Left)
	pathSumTraverse(root.Right)
	bTreePath = bTreePath[:len(bTreePath)-1]
}
