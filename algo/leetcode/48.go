/*
48. Rotate Image
Medium
16.6K
730
Companies
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
*/

package leetcode

// 顺时针旋转90度，可以先用左下角和右上角的对角线翻转一遍，然后再上下翻转一遍，即可得预期旋转效果
func Rotate90Degree(matrix [][]int) {
	n := len(matrix)
	// 沿对角线翻转
	for row := 0; row < n; row++ {
		for col := row; col < n; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}

	// 上下翻转
	for row := 0; row < n; row++ {
		topCol, bottomCol := 0, n-1
		for topCol < bottomCol {
			matrix[row][topCol], matrix[row][bottomCol] = matrix[row][bottomCol], matrix[row][topCol]
			topCol++
			bottomCol--
		}
	}
}
