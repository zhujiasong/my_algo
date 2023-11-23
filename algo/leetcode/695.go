/*
695. Max Area of Island
Medium
Topics
Companies
You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.



Example 1:


Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.
Example 2:

Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0
*/

package leetcode

func MaxAreaOfIsland(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	var maxArea, currArea int

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= row || j >= col {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		currArea++
		grid[i][j] = 0

		dfs(i, j+1)
		dfs(i, j-1)
		dfs(i-1, j)
		dfs(i+1, j)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				currArea = 0
				dfs(i, j)
				maxArea = max(maxArea, currArea)
			}
		}
	}

	return maxArea
}
