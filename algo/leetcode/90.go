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

var _subsetsWithDupRes [][]int
var _subsetWithDupTrack []int

func SubsetsWithDup(nums []int) [][]int {
	_subsetsWithDupRes = make([][]int, 0)
	_subsetWithDupTrack = make([]int, 0, len(nums))
	sli := sort.IntSlice(nums)
	sort.Sort(sli)

	_subsetWithDupBackTrack(nums, 0)

	return _subsetsWithDupRes
}

func _subsetWithDupBackTrack(nums []int, start int) {
	_subsetsWithDupRes = append(_subsetsWithDupRes, append([]int{}, _subsetWithDupTrack...))

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		_subsetWithDupTrack = append(_subsetWithDupTrack, nums[i])
		_subsetWithDupBackTrack(nums, i+1)
		_subsetWithDupTrack = _subsetWithDupTrack[:len(_subsetWithDupTrack)-1]
	}
}
