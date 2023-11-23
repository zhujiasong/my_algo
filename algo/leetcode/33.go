/*
33. Search in Rotated Sorted Array
Medium
24.9K
1.5K
Companies
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.
*/

package leetcode

func SearchRotatedArray(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, low, high, target int) int {
	for low <= high {
		mid := low + high/2 - 1
		if nums[mid] == target {
			return mid
		}
		if nums[low] == target {
			return low
		}
		if nums[high] == target {
			return high
		}

		if nums[low] < nums[mid] {
			if nums[mid] > target && target >= nums[low] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
