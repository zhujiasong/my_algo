package leetcode

func MaxSubArray(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}
	if l == 1 {
		return nums[0]
	}

	curSum := nums[0]
	maxSum := curSum

	for i := 1; i < l; i++ {
		if curSum <= 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}

		maxSum = max(maxSum, curSum)
	}

	return maxSum
}
