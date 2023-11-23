/*
40. Combination Sum II
Medium
9.9K
259
Companies
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.



Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

*/

package leetcode

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	l := len(candidates)
	res := make([][]int, 0)
	trackArr := make([]int, 0, l)
	sum := 0

	var backtrack func([]int, int)
	backtrack = func(nums []int, start int) {
		if sum == target {
			res = append(res, append([]int{}, trackArr...))
			return
		}
		if sum > target {
			return
		}

		for i := start; i < l; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			sum += nums[i]
			trackArr = append(trackArr, nums[i])

			backtrack(nums, i+1)

			sum -= nums[i]
			trackArr = trackArr[:len(trackArr)-1]
		}
	}

	sort.Ints(candidates)
	backtrack(candidates, 0)
	return res
}
