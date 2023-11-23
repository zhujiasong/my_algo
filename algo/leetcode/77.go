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

var _combineRes [][]int
var _combineTrack []int

func Combine(n, k int) [][]int {
	_combineRes = make([][]int, 0)
	_combineTrack = make([]int, 0)

	_combineBackTrack(n+1, 1, k)
	return _combineRes
}

func _combineBackTrack(n, start, k int) {
	if len(_combineTrack) == k {
		_combineRes = append(_combineRes, append([]int{}, _combineTrack...))
		return
	}

	for i := start; i < n; i++ {
		_combineTrack = append(_combineTrack, i)
		_combineBackTrack(n, i+1, k)
		_combineTrack = _combineTrack[:len(_combineTrack)-1]
	}
}
