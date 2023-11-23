/*
337. House Robber III
Medium
Topics
Companies
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root.

Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that all houses in this place form a binary tree. It will automatically contact the police if two directly-linked houses were broken into on the same night.

Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.



Example 1:


Input: root = [3,2,3,null,3,null,1]
Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
Example 2:


Input: root = [3,4,5,1,3,null,1]
Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.


Constraints:

The number of nodes in the tree is in the range [1, 104].
0 <= Node.val <= 104
*/

package leetcode

import (
	tree "my_algo/ds/binary_tree"
)

func rob3(root *tree.TreeNode) int {
	maximun := 0
	note := make(map[*tree.TreeNode]int)
	var dfs func(node *tree.TreeNode) int
	dfs = func(node *tree.TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return node.Val
		}

		var ll, lr, rl, rr int
		m := node.Val
		if node.Left != nil {
			if val, ok := note[node.Left.Left]; ok {
				ll = val
			} else {
				ll = dfs(node.Left.Left)
				note[node.Left.Left] = ll
			}
			if val, ok := note[node.Left.Right]; ok {
				lr = val
			} else {
				lr = dfs(node.Left.Right)
				note[node.Left.Right] = lr
			}
			m += ll + lr
		}
		if node.Right != nil {
			if val, ok := note[node.Right.Left]; ok {
				rl = val
			} else {
				rl = dfs(node.Right.Left)
				note[node.Right.Left] = rl
			}
			if val, ok := note[node.Right.Right]; ok {
				rr = val
			} else {
				rr = dfs(node.Right.Right)
				note[node.Right.Right] = rr
			}
			m += rl + rr
		}

		var l, r int
		if val, ok := note[node.Left]; ok {
			l = val
		} else {
			l = dfs(node.Left)
			note[node.Left] = l
		}
		if val, ok := note[node.Right]; ok {
			r = val
		} else {
			r = dfs(node.Right)
			note[node.Right] = r
		}
		m = max(m, l+r)

		maximun = max(maximun, m)

		return m
	}

	return max(dfs(root), maximun)
}
