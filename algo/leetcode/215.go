/*
215. Kth Largest Element in an Array
Medium
16.3K
803
Companies
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

Example 1:
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5

Example 2:
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4


Constraints:
*/

package leetcode

import (
	"my_algo/algo/sort"
)

func FindKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	return _findKthLargest(nums, k, 0, len(nums)-1)
}

func _findKthLargest(nums []int, k, low, high int) int {
	if low >= high {
		return nums[high]
	}

	index := sort.Partition(nums, low, high)
	len := len(nums) - index - 1
	if len == k-1 {
		return nums[index]
	} else if len > k-1 {
		low = index + 1
	} else {
		high = index - 1
	}
	return _findKthLargest(nums, k, low, high)
}
