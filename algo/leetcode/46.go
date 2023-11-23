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

func Permute(nums []int) [][]int {
	l := len(nums)
	res := make([][]int, 0)
	trackArr := make([]int, 0, l)
	visited := make([]bool, l)

	var backtrack func([]int)
	backtrack = func(nums []int) {
		if len(trackArr) == l {
			res = append(res, append([]int{}, trackArr...))
			return
		}

		for i := 0; i < l; i++ {
			if visited[i] {
				continue
			}

			trackArr = append(trackArr, nums[i])
			visited[i] = true

			backtrack(nums)

			visited[i] = false
			trackArr = trackArr[:len(trackArr)-1]
		}
	}

	backtrack(nums)
	return res
}
