/*
55. Jump Game
Attempted
Medium
Topics
Companies
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.



Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.


Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 105
*/

package leetcode

func CanJump(nums []int) bool {
	l := len(nums)
	dp := make([]bool, l)
	dp[0] = true

	for i := 1; i < l; i++ {
		j := i - 1
		for ; j >= 0; j-- {
			if dp[j] && i-j <= nums[j] {
				dp[i] = true
				if l-i-1 <= nums[i] {
					return true
				}
				break
			}
		}

	}

	return dp[l-1]
}

// copy leetcode
func CanJumpV2(nums []int) bool {
	maxJumpIdx := 0
	l := len(nums)
	for i := 0; i < l && maxJumpIdx >= i; i++ {
		if maxJumpIdx > l-i-1 {
			return true
		}
		maxJumpIdx = max(maxJumpIdx, i+nums[i])
	}

	return false
}
