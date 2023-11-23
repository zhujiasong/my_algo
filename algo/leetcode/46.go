/*
46. Permutations
Medium
18.2K
294
Companies
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.



Example 1:
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:
Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:
Input: nums = [1]
Output: [[1]]

*/

package leetcode

var _permuteRes [][]int
var _permuteTrack []int

func Permute(nums []int) [][]int {
	_permuteRes = make([][]int, 0)
	_permuteTrack = make([]int, 0)
	visited := make([]bool, len(nums))
	_permuteBackTrack(nums, visited)

	return _permuteRes
}

func _permuteBackTrack(nums []int, visited []bool) {
	if len(_permuteTrack) == len(nums) {
		_permuteRes = append(_permuteRes, append([]int{}, _permuteTrack...))
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}

		_permuteTrack = append(_permuteTrack, nums[i])
		visited[i] = true

		_permuteBackTrack(nums, visited)

		visited[i] = false
		_permuteTrack = _permuteTrack[:len(_permuteTrack)-1]
	}
}
