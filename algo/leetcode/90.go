/*
90. Subsets II
Medium
9.2K
265
Companies
Given an integer array nums that may contain duplicates, return all possible
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:
Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]

example 2:
Input: nums = [0]
Output: [[],[0]]

*/

package leetcode

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	l := len(nums)
	res := make([][]int, 0, l)
	trackArr := make([]int, 0, l)
	var backtrack func([]int, int)

	backtrack = func(nums []int, start int) {
		res = append(res, append([]int{}, trackArr...))

		for i := start; i < l; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			trackArr = append(trackArr, nums[i])
			backtrack(nums, i+1)
			trackArr = trackArr[:len(trackArr)-1]
		}
	}

	sort.Ints(nums)
	backtrack(nums, 0)

	return res
}
