/*
560. Subarray Sum Equals K
Attempted
Medium
Topics
Companies
Hint
Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

A subarray is a contiguous non-empty sequence of elements within an array.



Example 1:

Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2

[4, 2, 5, 1, 4, 9, 8] 5
[4, 6, 11,12,16,25,33]

Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107
*/

package leetcode

func SubarraySum(nums []int, k int) int {
	l := len(nums)
	prefixSum := make(map[int]int, l)
	prefixSum[0] = 1
	var sum, cnt int

	for i := range nums {
		sum += nums[i]

		if v, ok := prefixSum[sum-k]; ok {
			cnt += v
		}

		prefixSum[sum]++
	}

	return cnt
}
