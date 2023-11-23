/*
930. Binary Subarrays With Sum
Medium
Topics
Companies
Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.

A subarray is a contiguous part of the array.



Example 1:

Input: nums = [1,0,1,0,1], goal = 2
Output: 4
Explanation: The 4 subarrays are bolded and underlined below:
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
Example 2:

Input: nums = [0,0,0,0,0], goal = 0
Output: 15
*/

package leetcode

// TODO, not complete
func NumSubarraysWithSum(nums []int, goal int) int {
	var left, right, sum int
	var ret int

	for right < len(nums) {
		sum += nums[right]
		right++

		if sum == goal {
			for right < len(nums) && nums[right] == 0 {
				ret++
				right++
			}
		}

		for left < right && sum == goal {
			ret++
			sum -= nums[left]
			left++
		}
	}

	return ret
}
