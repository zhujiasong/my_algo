/*
39. Combination Sum
Medium
17.9K
367
Companies
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
frequency
 of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.



Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1

*/

package leetcode

func CombinationSum(candidates []int, target int) [][]int {
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
			sum += nums[i]
			trackArr = append(trackArr, nums[i])

			backtrack(nums, i)

			sum -= nums[i]
			trackArr = trackArr[:len(trackArr)-1]
		}
	}

	backtrack(candidates, 0)
	return res
}
