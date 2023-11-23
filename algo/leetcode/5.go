/*
5. Longest Palindromic Substring
Medium
Topics
Companies
Hint
Given a string s, return the longest
palindromic

substring
 in s.



Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"


Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters.
*/

package leetcode

func LongestPalindrome(s string) string {
	var start, end = -1, 0
	for i := range s {
		l1, r1 := palindrome(s, i, i)
		l2, r2 := palindrome(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}

	}
	return s[start+1 : end]
}

func palindrome(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}

	return left, right
}
