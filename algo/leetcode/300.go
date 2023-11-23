/*
Given an integer array nums, return the length of the longest strictly increasing
subsequence
.



Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1


Constraints:

1 <= nums.length <= 2500
-104 <= nums[i] <= 104
*/

package leetcode

func LengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	m := 0
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		m = max(dp[i], m)
	}

	return m
}

func MaxLIS(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	m := 0
	dp := make([]int, len(nums))
	flg := make([]int, len(nums))
	dp[0] = 1
	maxIdxs := make([]int, 1)

	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] <= nums[j] {
				continue
			}
			if dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
				// 关键,记录 上一个比nums[i]小的数中，拥有最大上升子序列的那个数的索引
				flg[i] = j
			}
		}
		if dp[i] > m {
			m = dp[i]
			maxIdxs[0] = i
			maxIdxs = maxIdxs[:1]
		} else if dp[i] == m {
			// 可能有多个最大长度
			maxIdxs = append(maxIdxs, i)
		}
	}

	out := make([][]int, len(maxIdxs))
	for i, maxIdx := range maxIdxs {
		out[i] = make([]int, m)
		outIdx := m - 1
		idx := maxIdx
		for n := 0; n < m; n++ {
			out[i][outIdx] = nums[idx]
			outIdx--
			idx = flg[idx]
		}
	}

	if len(out) == 1 {
		return out[0]
	}

	var ret []int

	// 查找最长序列中字典序最小的
	for col := 0; col < m; col++ {
		minNum := out[0][col]
		for row := 1; row < len(out); row++ {
			if out[row][col] < minNum {
				minNum = out[row][col]
				ret = out[row]
			}
		}
		if minNum != out[0][col] {
			return ret
		}
	}

	return out[0]
}
