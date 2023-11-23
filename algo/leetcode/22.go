/*
22. Generate Parentheses
Medium
Topics
Companies
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.



Example 1:

Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
Example 2:

Input: n = 1
Output: ["()"]


Constraints:

1 <= n <= 8
*/

package leetcode

func GenerateParenthesis(n int) []string {
	res := make([]string, 0)
	var currStr string

	var backtrack func(n, open, close int)
	backtrack = func(n, open, close int) {
		if open == close && open == n {
			res = append(res, currStr)
			return
		}

		if open < n {
			currStr += "("
			open++
			backtrack(n, open, close)
			open--
			currStr = currStr[:len(currStr)-1]
		}

		if close < open {
			currStr += ")"
			close++
			backtrack(n, open, close)
			close--
			currStr = currStr[:len(currStr)-1]
		}
	}

	backtrack(n, 0, 0)
	return res
}
