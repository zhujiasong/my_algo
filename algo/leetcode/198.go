/*
198. House Robber
Solved
Medium
Topics
Companies
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.



Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
*/

package leetcode

// fn = max(f(n-1), f(n-2)+nums[n])
func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	tmp1, tmp2, tmp3 := nums[0], max(nums[0], nums[1]), 0
	if l == 2 {
		return tmp2
	}

	for i := 2; i < l; i++ {
		tmp3 = max(tmp2, tmp1+nums[i])
		tmp1, tmp2 = tmp2, tmp3
	}

	return tmp3
}
