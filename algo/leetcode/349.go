/*
349. Intersection of Two Arrays
Attempted
Easy
Topics
Companies
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.



Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.


Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/

package leetcode

func Intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	n1, n2 := nums1, nums2
	if len(nums1) > len(nums2) {
		n1 = nums2
		n2 = nums1
	}
	m := make(map[int]int, len(n1))

	for _, num := range n1 {
		if _, ok := m[num]; !ok {
			m[num]++
		}
	}

	ret := make([]int, 0, len(n1))
	for _, num := range n2 {
		if _, ok := m[num]; ok {
			m[num]++
		}
	}
	for k, v := range m {
		if v > 1 {
			ret = append(ret, k)
		}
	}
	return ret
}
