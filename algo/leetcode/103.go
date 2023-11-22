package leetcode

import (
	tree "my_algo/ds/binary_tree"
)

func ZigzagLevelOrder(root *tree.TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*tree.TreeNode{root}
	ret := [][]int{{root.Val}}
	size := 1
	flg := false

	for len(queue) != 0 {
		tmpNode := queue[0]
		queue = queue[1:]
		size--
		if tmpNode.Left != nil {
			queue = append(queue, tmpNode.Left)
		}
		if tmpNode.Right != nil {
			queue = append(queue, tmpNode.Right)
		}
		if size == 0 && len(queue) != 0 {
			tmp := make([]int, 0, len(queue))
			if flg {
				for _, node := range queue {
					tmp = append(tmp, node.Val)
				}
			} else {
				for i := len(queue) - 1; i >= 0; i-- {
					tmp = append(tmp, queue[i].Val)
				}
			}
			ret = append(ret, tmp)
			size = len(queue)
			flg = !flg
		}
	}

	return ret
}
