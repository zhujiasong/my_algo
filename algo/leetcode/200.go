/*
200. Number of Islands
Medium
Topics
Companies
Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.

An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.



Example 1:

Input: grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
Output: 1
Example 2:

Input: grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Output: 3


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] is '0' or '1'.
*/

package leetcode

func NumIslands(grid [][]byte) int {
	var cnt int
	row := len(grid)
	col := len(grid[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfsNumIslands(grid, row, col, i, j)
			}
		}
	}

	return cnt
}

func dfsNumIslands(grid [][]byte, row, col, i, j int) {
	if i < 0 || j < 0 || i >= row || j >= col {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'

	dfsNumIslands(grid, row, col, i+1, j)
	dfsNumIslands(grid, row, col, i-1, j)
	dfsNumIslands(grid, row, col, i, j+1)
	dfsNumIslands(grid, row, col, i, j-1)
}
