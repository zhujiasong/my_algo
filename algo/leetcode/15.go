/*
15. 3Sum
Medium
Topics
Companies
Hint
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/

package leetcode

import "sort"

func ThreeSum(nums []int) [][]int {
	return threeSumK(nums, 0)
}

func threeSumK(nums []int, k int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	res := make([][]int, 0, 4)
	sort.Ints(nums)
	for i := 0; i <= len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		t := k - nums[i]
		for left < right {
			low, high := nums[left], nums[right]

			if t == nums[left]+nums[right] {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == low {
					left++
				}
				for left < right && nums[right] == high {
					right--
				}
			} else if t > nums[left]+nums[right] {
				for left < right && nums[left] == low {
					left++
				}
			} else {
				for left < right && nums[right] == high {
					right--
				}
			}
		}
	}

	return res
}
