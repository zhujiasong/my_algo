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

var _subsetsRes [][]int
var _subsetsTrack []int

func Subsets(nums []int) [][]int {
	_subsetsRes = make([][]int, 0)
	_subsetsTrack = make([]int, 0)
	_subsetsBackTrack(nums, 0)
	return _subsetsRes
}

func _subsetsBackTrack(nums []int, start int) {
	_subsetsRes = append(_subsetsRes, append([]int{}, _subsetsTrack...))

	for i := start; i < len(nums); i++ {
		_subsetsTrack = append(_subsetsTrack, nums[i])
		_subsetsBackTrack(nums, i+1)
		_subsetsTrack = _subsetsTrack[:len(_subsetsTrack)-1]
	}
}
