/*
78. Subsets
Medium
16.2K
240
Companies
Given an integer array nums of unique elements, return all possible
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [0]
Output: [[],[0]]
*/

package leetcode

func Subsets(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, n)
	trackArr := make([]int, 0, n)
	var backtrack func([]int, int)

	backtrack = func(nums []int, start int) {
		res = append(res, append([]int{}, trackArr...))

		for i := start; i < n; i++ {
			trackArr = append(trackArr, nums[i])
			backtrack(nums, i+1)
			trackArr = trackArr[:len(trackArr)-1]
		}
	}

	backtrack(nums, 0)

	return res
}
