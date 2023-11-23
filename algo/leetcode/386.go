/*
386. Lexicographical Numbers
Medium
Topics
Companies
Given an integer n, return all the numbers in the range [1, n] sorted in lexicographical order.

You must write an algorithm that runs in O(n) time and uses O(1) extra space.



Example 1:

Input: n = 13
Output: [1,10,11,12,13,2,3,4,5,6,7,8,9]
Example 2:

Input: n = 2
Output: [1,2]
*/

package leetcode

func LexicalOrder(n int) []int {
	ret := make([]int, n)
	number := 1

	for i := 0; i < n; i++ {
		ret[i] = number
		if number*10 <= n {
			number *= 10
		} else {
			for number%10 == 9 || number+1 > n {
				number /= 10
			}
			number++
		}
	}

	return ret
}
