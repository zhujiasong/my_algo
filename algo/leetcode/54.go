/*
54. Spiral Matrix
Medium
13.9K
1.2K
Companies
Given an m x n matrix, return all elements of the matrix in spiral order.

*/

package leetcode

func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	// right, down, left, up
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d := 0

	rowBegin, rowEnd, colBegin, colEnd := 0, m-1, 0, n-1
	row, col := 0, 0

	res := make([]int, 0, m*n)
	for {
		for row >= rowBegin && row <= rowEnd && col >= colBegin && col <= colEnd {
			res = append(res, matrix[row][col])
			row += direction[d][0]
			col += direction[d][1]
		}

		if len(res) >= m*n {
			break
		}

		switch d % 4 {
		case 0:
			rowBegin++
			row++
			col--
		case 1:
			colEnd--
			row--
			col--
		case 2:
			rowEnd--
			row--
			col++
		case 3:
			colBegin++
			row++
			col++
		}

		d++
		if d == 4 {
			d = 0
		}
	}

	return res
}
