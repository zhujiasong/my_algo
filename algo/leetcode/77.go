/*
77. Combinations
Medium
7.9K
206
Companies
Given two integers n and k, return all possible combinations of k numbers chosen from the range [1, n].

You may return the answer in any order.



Example 1:

Input: n = 4, k = 2
Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
Explanation: There are 4 choose 2 = 6 total combinations.
Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to be the same combination.
Example 2:

Input: n = 1, k = 1
Output: [[1]]
Explanation: There is 1 choose 1 = 1 total combination.
*/

package leetcode

func Combine(n int, k int) [][]int {
	res := make([][]int, 0)
	if n < k {
		return [][]int{}
	}

	trackArr := make([]int, 0, n)
	var backtrack func(int, int)

	backtrack = func(n, start int) {
		if len(trackArr) == k {
			res = append(res, append([]int{}, trackArr...))
			return
		}

		for i := start; i <= n; i++ {
			trackArr = append(trackArr, i)
			backtrack(n, i+1)
			trackArr = trackArr[:len(trackArr)-1]
		}
	}

	backtrack(n, 1)

	return res
}
