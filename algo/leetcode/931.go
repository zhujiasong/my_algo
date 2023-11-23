/*
931. Minimum Falling Path Sum
Medium
5.4K
132
Companies
Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.

A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right. Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).
*/

package leetcode

import "math"

func MinFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	minPathSum := math.MaxInt
	rowNum, colNum := len(matrix), len(matrix[0])
	prePathSum := make([]int, 0, colNum)
	currPathSum := make([]int, colNum)

	prePathSum = append(prePathSum, matrix[0]...)

	for row := 1; row < rowNum; row++ {
		for col := 0; col < colNum; col++ {
			currPathSum[col] = matrix[row][col] + minPrevRowNeighbor(prePathSum, col)
			if row == rowNum-1 {
				minPathSum = min(minPathSum, currPathSum[col])
			}
		}
		prePathSum = append([]int{}, currPathSum...)
	}

	return minPathSum
}

func minPrevRowNeighbor(nums []int, col int) int {
	m := nums[col]
	if col > 0 {
		m = min(m, nums[col-1])
	}
	if col < len(nums)-1 {
		m = min(m, nums[col+1])
	}
	return m
}
