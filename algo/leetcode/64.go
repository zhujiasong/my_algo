/*
64. Minimum Path Sum
Solved
Medium
Topics
Companies
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.



Example 1:


Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
Example 2:

Input: grid = [[1,2,3],[4,5,6]]
Output: 12
*/

package leetcode

func MinPathSum(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	dp := make([]int, col)
	dp[0] = grid[0][0]
	for i := 1; i < col; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}

	for i := 1; i < row; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < col; j++ {
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}

	return dp[len(dp)-1]
}
