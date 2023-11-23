/*
62. Unique Paths
Medium
Topics
Companies
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.
*/

package leetcode

func UniquePaths(m int, n int) int {
	if n == 1 || m == 1 {
		return 1
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for ; m > 1; m-- {
		for i := 1; i < n; i++ {
			dp[i] += dp[i-1]
		}
	}
	return dp[len(dp)-1]
}
