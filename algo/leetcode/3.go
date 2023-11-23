/*
Given a string s, find the length of the longest
substring
 without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

package leetcode

func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	maxLen := 0
	window := make(map[byte]int, len(s))
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		window[c]++
		right++

		for left < right && window[c] > 1 {
			window[s[left]]--
			left++
		}

		if right-left > maxLen {
			maxLen = right - left
		}
	}

	return maxLen
}
