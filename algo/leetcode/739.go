/*
739. Daily Temperatures
Medium
11.9K
262
Companies
Given an array of integers temperatures represents the daily temperatures, return an array answer such that answer[i] is the number of days you have to wait after the ith day to get a warmer temperature. If there is no future day for which this is possible, keep answer[i] == 0 instead.



Example 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
Example 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
Example 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]


Constraints:

1 <= temperatures.length <= 105
30 <= temperatures[i] <= 100
*/

package leetcode

func DailyTemperatures(temperatures []int) []int {
	if len(temperatures) <= 1 {
		return []int{0}
	}

	stack := make([]int, len(temperatures))
	top := -1
	res := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		for top != -1 && temperatures[stack[top]] <= temperatures[i] {
			top--
		}

		res[i] = 0
		if top != -1 {
			res[i] = stack[top] - i
		}
		top++
		stack[top] = i
	}

	return res
}
