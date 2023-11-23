/*
124. Binary Tree Maximum Path Sum
Hard
Topics
Companies
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.



Example 1:


Input: root = [1,2,3]
Output: 6
Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
Example 2:


Input: root = [-10,9,20,null,null,15,7]
Output: 42
Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.


Constraints:

The number of nodes in the tree is in the range [1, 3 * 104].
-1000 <= Node.val <= 1000
*/

package leetcode

import (
	"math"
	tree "my_algo/ds/binary_tree"
)

func MaxPathSum(root *tree.TreeNode) int {
	maxPath := math.MinInt
	var dfs func(*tree.TreeNode) int

	dfs = func(root *tree.TreeNode) int {
		if root == nil {
			return 0
		}

		l := dfs(root.Left)
		r := dfs(root.Right)

		sum := root.Val
		if l > 0 {
			sum += l
		}
		if r > 0 {
			sum += r
		}
		maxPath = max(maxPath, sum)

		return max(max(l, r)+root.Val, root.Val)
	}

	dfs(root)
	return maxPath
}
