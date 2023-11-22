/*
76. Minimum Window Substring
Hard
16.5K
678
Companies
Given two strings s and t of lengths m and n respectively, return the minimum window
substring
 of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
*/
package leetcode

import "math"

func MinWindow(s string, t string) string {
	var ret string
	var min = math.MaxInt
	var vaild int
	windowT := make(map[byte]int, len(t))
	for _, b := range t {
		windowT[byte(b)]++
	}

	windowS := make(map[byte]int, len(t))
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := windowT[c]; ok {
			windowS[c]++
			if windowS[c] == windowT[c] {
				vaild++
			}
		}

		for left < right {
			if vaild != len(windowT) {
				break
			}
			if right-left < min {
				min = right - left
				ret = s[left:right]
			}

			d := s[left]
			left++
			if _, ok := windowT[d]; ok {
				if windowS[d] == windowT[d] {
					vaild--
				}
				windowS[d]--
			}
			if left < len(s) && windowS[s[left]] < windowT[s[left]] {
				break
			}
		}
	}

	return ret
}
