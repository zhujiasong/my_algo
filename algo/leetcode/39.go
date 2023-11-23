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

var _combinationSumRes [][]int
var _combinationSumTrack []int

func CombinationSum(candidates []int, target int) [][]int {
	_combinationSumRes = make([][]int, 0)
	_combinationSumTrack = make([]int, 0, len(candidates))

	_combinationSumBackTrack(candidates, 0, target, 0)

	return _combinationSumRes
}

func _combinationSumBackTrack(candidates []int, start, target, currSum int) {
	if currSum == target {
		_combinationSumRes = append(_combinationSumRes, append([]int{}, _combinationSumTrack...))
		return
	}
	if currSum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		_combinationSumTrack = append(_combinationSumTrack, candidates[i])
		currSum += candidates[i]

		_combinationSumBackTrack(candidates, i, target, currSum)

		lastIdx := len(_combinationSumTrack) - 1
		currSum -= _combinationSumTrack[lastIdx]
		_combinationSumTrack = _combinationSumTrack[:lastIdx]
	}
}
