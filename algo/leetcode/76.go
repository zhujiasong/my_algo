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
	ls, lt := len(s), len(t)
	if ls < lt {
		return ""
	}

	windowS := make(map[byte]int, ls)
	windowT := make(map[byte]int, lt)
	vaild := 0
	left, right := 0, 0
	start, distance := 0, math.MaxInt

	for i := range t {
		windowT[t[i]]++
	}

	for right < ls {
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

			d := s[left]
			if _, ok := windowT[d]; ok {
				if windowS[d] == windowT[d] {
					if right-left < distance {
						start = left
						distance = right - left
					}
					vaild--
				}
				windowS[d]--
			}
			left++
		}
	}

	if distance == math.MaxInt {
		return ""
	}
	return s[start : start+distance]
}
