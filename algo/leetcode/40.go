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

var _combinationSum2Res [][]int
var _combinationSum2Track []int

func CombinationSum2(candidates []int, target int) [][]int {
	_combinationSum2Res = make([][]int, 0)
	_combinationSum2Track = make([]int, 0, len(candidates))
	sli := sort.IntSlice(candidates)
	sort.Sort(sli)

	_combinationSum2BackTrack(candidates, 0, target, 0)
	return _combinationSum2Res
}

func _combinationSum2BackTrack(candidates []int, start, target, currSum int) {
	if target == currSum {
		_combinationSum2Res = append(_combinationSum2Res, append([]int{}, _combinationSum2Track...))
		return
	}

	if currSum > target {
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		_combinationSum2Track = append(_combinationSum2Track, candidates[i])
		currSum += candidates[i]

		_combinationSum2BackTrack(candidates, i+1, target, currSum)

		lastIdx := len(_combinationSum2Track) - 1
		currSum -= _combinationSum2Track[lastIdx]
		_combinationSum2Track = _combinationSum2Track[:lastIdx]
	}
}
