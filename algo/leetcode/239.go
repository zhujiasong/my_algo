/*
You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.



Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Example 2:

Input: nums = [1], k = 1
Output: [1]
*/

package leetcode

import "container/list"

func MaxSlidingWindow(nums []int, k int) []int {
	queue := list.New()
	var left, right int
	ret := make([]int, 0, len(nums)-k+1)

	for right < len(nums) {
		num := nums[right]
		right++
		for queue.Len() > 0 && num > queue.Back().Value.(int) {
			queue.Remove(queue.Back())
		}
		queue.PushBack(num)

		if right < k {
			continue
		}

		currMax := queue.Front().Value.(int)
		ret = append(ret, currMax)

		if nums[left] == currMax {
			queue.Remove(queue.Front())
		}
		left++
	}

	return ret
}
