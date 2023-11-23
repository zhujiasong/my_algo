/*
496. Next Greater Element I
Easy
7.3K
578
Companies
The next greater element of some element x in an array is the first greater element that is to the right of x in the same array.

You are given two distinct 0-indexed integer arrays nums1 and nums2, where nums1 is a subset of nums2.

For each 0 <= i < nums1.length, find the index j such that nums1[i] == nums2[j] and determine the next greater element of nums2[j] in nums2. If there is no next greater element, then the answer for this query is -1.

Return an array ans of length nums1.length such that ans[i] is the next greater element as described above.



Example 1:

Input: nums1 = [4,1,2], nums2 = [1,3,4,2]
Output: [-1,3,-1]
Explanation: The next greater element for each value of nums1 is as follows:
- 4 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
- 1 is underlined in nums2 = [1,3,4,2]. The next greater element is 3.
- 2 is underlined in nums2 = [1,3,4,2]. There is no next greater element, so the answer is -1.
Example 2:

Input: nums1 = [2,4], nums2 = [1,2,3,4]
Output: [3,-1]
Explanation: The next greater element for each value of nums1 is as follows:
- 2 is underlined in nums2 = [1,2,3,4]. The next greater element is 3.
- 4 is underlined in nums2 = [1,2,3,4]. There is no next greater element, so the answer is -1.
*/

package leetcode

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	num1IndexMap := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		num1IndexMap[v] = -1
	}
	for i, v := range nums2 {
		if _, ok := num1IndexMap[v]; ok {
			num1IndexMap[v] = i
		}
	}

	resMap := make(map[int]int, len(nums1))
	stack := make([]int, len(nums2))
	top := -1
	for i := len(nums2) - 1; i >= 0; i-- {
		for top != -1 && stack[top] <= nums2[i] {
			top--
		}

		if idx, ok := num1IndexMap[nums2[i]]; ok && idx == i {
			resMap[i] = -1
			if top != -1 {
				resMap[i] = stack[top]
			}
		}

		top++
		stack[top] = nums2[i]
	}

	res := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		idx := num1IndexMap[v]
		res = append(res, resMap[idx])
	}

	return res
}
