/*
912. Sort an Array
Medium
5.7K
724
Companies
Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.
*/

package leetcode

import (
	"my_algo/algo/sort"
)

func SortArray(nums []int) []int {
	return sort.QuickSort(nums)
}
